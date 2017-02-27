/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package deployment contains all the logic for handling Kubernetes Deployments.
// It implements a set of strategies (rolling, recreate) for deploying an application,
// the means to rollback to previous versions, proportional scaling for mitigating
// risk, cleanup policy, and other useful features of Deployments.
package deployment

import (
	"fmt"
	"reflect"
	"time"

	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	clientv1 "k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/v1"
	extensions "k8s.io/kubernetes/pkg/apis/extensions/v1beta1"
	"k8s.io/kubernetes/pkg/client/clientset_generated/clientset"
	coreinformers "k8s.io/kubernetes/pkg/client/informers/informers_generated/externalversions/core/v1"
	extensionsinformers "k8s.io/kubernetes/pkg/client/informers/informers_generated/externalversions/extensions/v1beta1"
	corelisters "k8s.io/kubernetes/pkg/client/listers/core/v1"
	extensionslisters "k8s.io/kubernetes/pkg/client/listers/extensions/v1beta1"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/controller/deployment/util"
	"k8s.io/kubernetes/pkg/util/metrics"
)

const (
	// maxRetries is the number of times a deployment will be retried before it is dropped out of the queue.
	maxRetries = 5
)

// controllerKind contains the schema.GroupVersionKind for this controller type.
var controllerKind = extensions.SchemeGroupVersion.WithKind("Deployment")

// DeploymentController is responsible for synchronizing Deployment objects stored
// in the system with actual running replica sets and pods.
type DeploymentController struct {
	// rsControl is used for adopting/releasing replica sets.
	rsControl     controller.RSControlInterface
	client        clientset.Interface
	eventRecorder record.EventRecorder

	// To allow injection of syncDeployment for testing.
	syncHandler func(dKey string) error
	// used for unit testing
	enqueueDeployment func(deployment *extensions.Deployment)

	// dLister can list/get deployments from the shared informer's store
	dLister extensionslisters.DeploymentLister
	// rsLister can list/get replica sets from the shared informer's store
	rsLister extensionslisters.ReplicaSetLister
	// podLister can list/get pods from the shared informer's store
	podLister corelisters.PodLister

	// dListerSynced returns true if the Deployment store has been synced at least once.
	// Added as a member to the struct to allow injection for testing.
	dListerSynced cache.InformerSynced
	// rsListerSynced returns true if the ReplicaSet store has been synced at least once.
	// Added as a member to the struct to allow injection for testing.
	rsListerSynced cache.InformerSynced
	// podListerSynced returns true if the pod store has been synced at least once.
	// Added as a member to the struct to allow injection for testing.
	podListerSynced cache.InformerSynced

	// Deployments that need to be synced
	queue workqueue.RateLimitingInterface
}

// NewDeploymentController creates a new DeploymentController.
func NewDeploymentController(dInformer extensionsinformers.DeploymentInformer, rsInformer extensionsinformers.ReplicaSetInformer, podInformer coreinformers.PodInformer, client clientset.Interface) *DeploymentController {
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	// TODO: remove the wrapper when every clients have moved to use the clientset.
	eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: v1core.New(client.Core().RESTClient()).Events("")})

	if client != nil && client.Core().RESTClient().GetRateLimiter() != nil {
		metrics.RegisterMetricAndTrackRateLimiterUsage("deployment_controller", client.Core().RESTClient().GetRateLimiter())
	}
	dc := &DeploymentController{
		client:        client,
		eventRecorder: eventBroadcaster.NewRecorder(api.Scheme, clientv1.EventSource{Component: "deployment-controller"}),
		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "deployment"),
	}
	dc.rsControl = controller.RealRSControl{
		KubeClient: client,
		Recorder:   dc.eventRecorder,
	}

	dInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    dc.addDeployment,
		UpdateFunc: dc.updateDeployment,
		// This will enter the sync loop and no-op, because the deployment has been deleted from the store.
		DeleteFunc: dc.deleteDeployment,
	})
	rsInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    dc.addReplicaSet,
		UpdateFunc: dc.updateReplicaSet,
		DeleteFunc: dc.deleteReplicaSet,
	})
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: dc.deletePod,
	})

	dc.syncHandler = dc.syncDeployment
	dc.enqueueDeployment = dc.enqueue

	dc.dLister = dInformer.Lister()
	dc.rsLister = rsInformer.Lister()
	dc.podLister = podInformer.Lister()
	dc.dListerSynced = dInformer.Informer().HasSynced
	dc.rsListerSynced = rsInformer.Informer().HasSynced
	dc.podListerSynced = podInformer.Informer().HasSynced
	return dc
}

// Run begins watching and syncing.
func (dc *DeploymentController) Run(workers int, stopCh <-chan struct{}) {
	defer utilruntime.HandleCrash()
	defer dc.queue.ShutDown()

	glog.Infof("Starting deployment controller")

	if !cache.WaitForCacheSync(stopCh, dc.dListerSynced, dc.rsListerSynced, dc.podListerSynced) {
		utilruntime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
		return
	}

	for i := 0; i < workers; i++ {
		go wait.Until(dc.worker, time.Second, stopCh)
	}

	<-stopCh
	glog.Infof("Shutting down deployment controller")
}

func (dc *DeploymentController) addDeployment(obj interface{}) {
	d := obj.(*extensions.Deployment)
	glog.V(4).Infof("Adding deployment %s", d.Name)
	dc.enqueueDeployment(d)
}

func (dc *DeploymentController) updateDeployment(old, cur interface{}) {
	oldD := old.(*extensions.Deployment)
	curD := cur.(*extensions.Deployment)
	glog.V(4).Infof("Updating deployment %s", oldD.Name)
	dc.enqueueDeployment(curD)
}

func (dc *DeploymentController) deleteDeployment(obj interface{}) {
	d, ok := obj.(*extensions.Deployment)
	if !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Couldn't get object from tombstone %#v", obj))
			return
		}
		d, ok = tombstone.Obj.(*extensions.Deployment)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Tombstone contained object that is not a Deployment %#v", obj))
			return
		}
	}
	glog.V(4).Infof("Deleting deployment %s", d.Name)
	dc.enqueueDeployment(d)
}

// addReplicaSet enqueues the deployment that manages a ReplicaSet when the ReplicaSet is created.
func (dc *DeploymentController) addReplicaSet(obj interface{}) {
	rs := obj.(*extensions.ReplicaSet)
	glog.V(4).Infof("ReplicaSet %s added.", rs.Name)

	if rs.DeletionTimestamp != nil {
		// On a restart of the controller manager, it's possible for an object to
		// show up in a state that is already pending deletion.
		dc.deleteReplicaSet(rs)
		return
	}

	// If it has a ControllerRef, that's all that matters.
	if controllerRef := controller.GetControllerOf(rs); controllerRef != nil {
		if controllerRef.Kind != controllerKind.Kind {
			// It's controller by a different type of controller.
			return
		}
		d, err := dc.dLister.Deployments(rs.Namespace).Get(controllerRef.Name)
		if err != nil {
			return
		}
		dc.enqueueDeployment(d)
		return
	}

	// Otherwise, it's an orphan. Get a list of all matching Deployments and sync
	// them to see if anyone wants to adopt it.
	for _, d := range dc.getDeploymentsForReplicaSet(rs) {
		dc.enqueueDeployment(d)
	}
}

// getDeploymentsForReplicaSet returns a list of Deployments that potentially
// match a ReplicaSet.
func (dc *DeploymentController) getDeploymentsForReplicaSet(rs *extensions.ReplicaSet) []*extensions.Deployment {
	deployments, err := dc.dLister.GetDeploymentsForReplicaSet(rs)
	if err != nil || len(deployments) == 0 {
		glog.V(4).Infof("Error: %v. No deployment found for ReplicaSet %v, deployment controller will avoid syncing.", err, rs.Name)
		return nil
	}
	// Because all ReplicaSet's belonging to a deployment should have a unique label key,
	// there should never be more than one deployment returned by the above method.
	// If that happens we should probably dynamically repair the situation by ultimately
	// trying to clean up one of the controllers, for now we just return the older one
	if len(deployments) > 1 {
		// ControllerRef will ensure we don't do anything crazy, but more than one
		// item in this list nevertheless constitutes user error.
		glog.V(4).Infof("user error! more than one deployment is selecting replica set %s/%s with labels: %#v, returning %s/%s",
			rs.Namespace, rs.Name, rs.Labels, deployments[0].Namespace, deployments[0].Name)
	}
	return deployments
}

// updateReplicaSet figures out what deployment(s) manage a ReplicaSet when the ReplicaSet
// is updated and wake them up. If the anything of the ReplicaSets have changed, we need to
// awaken both the old and new deployments. old and cur must be *extensions.ReplicaSet
// types.
func (dc *DeploymentController) updateReplicaSet(old, cur interface{}) {
	curRS := cur.(*extensions.ReplicaSet)
	oldRS := old.(*extensions.ReplicaSet)
	if curRS.ResourceVersion == oldRS.ResourceVersion {
		// Periodic resync will send update events for all known replica sets.
		// Two different versions of the same replica set will always have different RVs.
		return
	}
	glog.V(4).Infof("ReplicaSet %s updated.", curRS.Name)

	labelChanged := !reflect.DeepEqual(curRS.Labels, oldRS.Labels)

	curControllerRef := controller.GetControllerOf(curRS)
	oldControllerRef := controller.GetControllerOf(oldRS)
	controllerRefChanged := !reflect.DeepEqual(curControllerRef, oldControllerRef)
	if controllerRefChanged &&
		oldControllerRef != nil && oldControllerRef.Kind == controllerKind.Kind {
		// The ControllerRef was changed. Sync the old controller, if any.
		d, err := dc.dLister.Deployments(oldRS.Namespace).Get(oldControllerRef.Name)
		if err == nil {
			dc.enqueueDeployment(d)
		}
	}

	// If it has a ControllerRef, that's all that matters.
	if curControllerRef != nil {
		if curControllerRef.Kind != controllerKind.Kind {
			// It's controlled by a different type of controller.
			return
		}
		d, err := dc.dLister.Deployments(curRS.Namespace).Get(curControllerRef.Name)
		if err != nil {
			return
		}
		dc.enqueueDeployment(d)
		return
	}

	// Otherwise, it's an orphan. If anything changed, sync matching controllers
	// to see if anyone wants to adopt it now.
	if labelChanged || controllerRefChanged {
		for _, d := range dc.getDeploymentsForReplicaSet(curRS) {
			dc.enqueueDeployment(d)
		}
	}
}

// deleteReplicaSet enqueues the deployment that manages a ReplicaSet when
// the ReplicaSet is deleted. obj could be an *extensions.ReplicaSet, or
// a DeletionFinalStateUnknown marker item.
func (dc *DeploymentController) deleteReplicaSet(obj interface{}) {
	rs, ok := obj.(*extensions.ReplicaSet)

	// When a delete is dropped, the relist will notice a pod in the store not
	// in the list, leading to the insertion of a tombstone object which contains
	// the deleted key/value. Note that this value might be stale. If the ReplicaSet
	// changed labels the new deployment will not be woken up till the periodic resync.
	if !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Couldn't get object from tombstone %#v", obj))
			return
		}
		rs, ok = tombstone.Obj.(*extensions.ReplicaSet)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Tombstone contained object that is not a ReplicaSet %#v", obj))
			return
		}
	}
	glog.V(4).Infof("ReplicaSet %s deleted.", rs.Name)

	controllerRef := controller.GetControllerOf(rs)
	if controllerRef == nil {
		// No controller should care about orphans being deleted.
		return
	}
	if controllerRef.Kind != controllerKind.Kind {
		// It's controlled by a different type of controller.
		return
	}

	d, err := dc.dLister.Deployments(rs.Namespace).Get(controllerRef.Name)
	if err != nil {
		return
	}
	dc.enqueueDeployment(d)
}

// deletePod will enqueue a Recreate Deployment once all of its pods have stopped running.
func (dc *DeploymentController) deletePod(obj interface{}) {
	pod, ok := obj.(*v1.Pod)

	// When a delete is dropped, the relist will notice a pod in the store not
	// in the list, leading to the insertion of a tombstone object which contains
	// the deleted key/value. Note that this value might be stale. If the Pod
	// changed labels the new deployment will not be woken up till the periodic resync.
	if !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Couldn't get object from tombstone %#v", obj))
			return
		}
		pod, ok = tombstone.Obj.(*v1.Pod)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Tombstone contained object that is not a pod %#v", obj))
			return
		}
	}
	if d := dc.getDeploymentForPod(pod); d != nil && d.Spec.Strategy.Type == extensions.RecreateDeploymentStrategyType {
		// Sync if this Deployment now has no more Pods.
		rsList, err := dc.getReplicaSetsForDeployment(d)
		if err != nil {
			return
		}
		podMap, err := dc.getPodMapForReplicaSets(d.Namespace, rsList)
		if err != nil {
			return
		}
		numPods := 0
		for _, podList := range podMap {
			numPods += len(podList.Items)
		}
		if numPods == 0 {
			dc.enqueueDeployment(d)
		}
	}
}

func (dc *DeploymentController) enqueue(deployment *extensions.Deployment) {
	key, err := controller.KeyFunc(deployment)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %#v: %v", deployment, err))
		return
	}

	dc.queue.Add(key)
}

func (dc *DeploymentController) enqueueRateLimited(deployment *extensions.Deployment) {
	key, err := controller.KeyFunc(deployment)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %#v: %v", deployment, err))
		return
	}

	dc.queue.AddRateLimited(key)
}

// enqueueAfter will enqueue a deployment after the provided amount of time.
func (dc *DeploymentController) enqueueAfter(deployment *extensions.Deployment, after time.Duration) {
	key, err := controller.KeyFunc(deployment)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %#v: %v", deployment, err))
		return
	}

	dc.queue.AddAfter(key, after)
}

// getDeploymentForPod returns the deployment managing the given Pod.
func (dc *DeploymentController) getDeploymentForPod(pod *v1.Pod) *extensions.Deployment {
	// Find the owning replica set
	var rs *extensions.ReplicaSet
	var err error
	controllerRef := controller.GetControllerOf(pod)
	if controllerRef == nil {
		// No controller owns this Pod.
		return nil
	}
	if controllerRef.Kind != extensions.SchemeGroupVersion.WithKind("ReplicaSet").Kind {
		// Not a pod owned by a replica set.
		return nil
	}
	rs, err = dc.rsLister.ReplicaSets(pod.Namespace).Get(controllerRef.Name)
	if err != nil {
		glog.V(4).Infof("Cannot get replicaset %q for pod %q: %v", controllerRef.Name, pod.Name, err)
		return nil
	}

	// Now find the Deployment that owns that ReplicaSet.
	controllerRef = controller.GetControllerOf(rs)
	if controllerRef == nil {
		return nil
	}
	if controllerRef.Kind != controllerKind.Kind {
		return nil
	}
	d, err := dc.dLister.Deployments(rs.Namespace).Get(controllerRef.Name)
	if err != nil {
		return nil
	}
	return d
}

// worker runs a worker thread that just dequeues items, processes them, and marks them done.
// It enforces that the syncHandler is never invoked concurrently with the same key.
func (dc *DeploymentController) worker() {
	for dc.processNextWorkItem() {
	}
}

func (dc *DeploymentController) processNextWorkItem() bool {
	key, quit := dc.queue.Get()
	if quit {
		return false
	}
	defer dc.queue.Done(key)

	err := dc.syncHandler(key.(string))
	dc.handleErr(err, key)

	return true
}

func (dc *DeploymentController) handleErr(err error, key interface{}) {
	if err == nil {
		dc.queue.Forget(key)
		return
	}

	if dc.queue.NumRequeues(key) < maxRetries {
		glog.V(2).Infof("Error syncing deployment %v: %v", key, err)
		dc.queue.AddRateLimited(key)
		return
	}

	utilruntime.HandleError(err)
	glog.V(2).Infof("Dropping deployment %q out of the queue: %v", key, err)
	dc.queue.Forget(key)
}

// getReplicaSetsForDeployment uses ControllerRefManager to reconcile
// ControllerRef by adopting and orphaning.
// It returns the list of ReplicaSets that this Deployment should manage.
func (dc *DeploymentController) getReplicaSetsForDeployment(d *extensions.Deployment) ([]*extensions.ReplicaSet, error) {
	rsList, err := dc.rsLister.ReplicaSets(d.Namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}
	deploymentSelector, err := metav1.LabelSelectorAsSelector(d.Spec.Selector)
	if err != nil {
		return nil, fmt.Errorf("deployment %s/%s has invalid label selector: %v", d.Namespace, d.Name, err)
	}
	cm := controller.NewReplicaSetControllerRefManager(dc.rsControl, d, deploymentSelector, controllerKind)
	return cm.ClaimReplicaSets(rsList)
}

// getPodMapForReplicaSets scans the list of all Pods and returns a map from
// RS UID to Pods controlled by that RS, based on the Pod's ControllerRef.
func (dc *DeploymentController) getPodMapForReplicaSets(namespace string, rsList []*extensions.ReplicaSet) (map[types.UID]*v1.PodList, error) {
	// List all Pods.
	pods, err := dc.podLister.Pods(namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}
	// Group Pods by their controller (if it's in rsList).
	podMap := make(map[types.UID]*v1.PodList, len(rsList))
	for _, rs := range rsList {
		podMap[rs.UID] = &v1.PodList{}
	}
	for _, pod := range pods {
		// Ignore inactive Pods since that's what ReplicaSet does.
		if !controller.IsPodActive(pod) {
			continue
		}
		controllerRef := controller.GetControllerOf(pod)
		if controllerRef == nil {
			continue
		}
		// Only append if we care about this UID.
		if podList, ok := podMap[controllerRef.UID]; ok {
			podList.Items = append(podList.Items, *pod)
		}
	}
	return podMap, nil
}

// syncDeployment will sync the deployment with the given key.
// This function is not meant to be invoked concurrently with the same key.
func (dc *DeploymentController) syncDeployment(key string) error {
	startTime := time.Now()
	glog.V(4).Infof("Started syncing deployment %q (%v)", key, startTime)
	defer func() {
		glog.V(4).Infof("Finished syncing deployment %q (%v)", key, time.Now().Sub(startTime))
	}()

	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	deployment, err := dc.dLister.Deployments(namespace).Get(name)
	if errors.IsNotFound(err) {
		glog.Infof("Deployment has been deleted %v", key)
		return nil
	}
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Unable to retrieve deployment %v from store: %v", key, err))
		return err
	}

	// Deep-copy otherwise we are mutating our cache.
	// TODO: Deep-copy only when needed.
	d, err := util.DeploymentDeepCopy(deployment)
	if err != nil {
		return err
	}

	everything := metav1.LabelSelector{}
	if reflect.DeepEqual(d.Spec.Selector, &everything) {
		dc.eventRecorder.Eventf(d, v1.EventTypeWarning, "SelectingAll", "This deployment is selecting all pods. A non-empty selector is required.")
		if d.Status.ObservedGeneration < d.Generation {
			d.Status.ObservedGeneration = d.Generation
			dc.client.Extensions().Deployments(d.Namespace).UpdateStatus(d)
		}
		return nil
	}

	// List ReplicaSets owned by this Deployment, while reconciling ControllerRef
	// through adoption/orphaning.
	rsList, err := dc.getReplicaSetsForDeployment(d)
	if err != nil {
		return err
	}
	// List all Pods owned by this Deployment, grouped by their ReplicaSet.
	// This is expensive, so do it once and pass it along to subroutines.
	podMap, err := dc.getPodMapForReplicaSets(d.Namespace, rsList)
	if err != nil {
		return err
	}

	if d.DeletionTimestamp != nil {
		return dc.syncStatusOnly(d, rsList, podMap)
	}

	// Why run the cleanup policy only when there is no rollback request?
	// The thing with the cleanup policy currently is that it is far from smart because it takes into account
	// the latest replica sets while it should instead retain the latest *working* replica sets. This means that
	// you can have a cleanup policy of 1 but your last known working replica set may be 2 or 3 versions back
	// in the history.
	// Eventually we will want to find a way to recognize replica sets that have worked at some point in time
	// (and chances are higher that they will work again as opposed to others that didn't) for candidates to
	// automatically roll back to (#23211) and the cleanup policy should help.
	if d.Spec.RollbackTo == nil {
		_, oldRSs, err := dc.getAllReplicaSetsAndSyncRevision(d, rsList, podMap, false)
		if err != nil {
			return err
		}
		// So far the cleanup policy was executed once a deployment was paused, scaled up/down, or it
		// succesfully completed deploying a replica set. Decouple it from the strategies and have it
		// run almost unconditionally - cleanupDeployment is safe by default.
		dc.cleanupDeployment(oldRSs, d)
	}

	// Update deployment conditions with an Unknown condition when pausing/resuming
	// a deployment. In this way, we can be sure that we won't timeout when a user
	// resumes a Deployment with a set progressDeadlineSeconds.
	if err = dc.checkPausedConditions(d); err != nil {
		return err
	}

	_, err = dc.hasFailed(d, rsList, podMap)
	if err != nil {
		return err
	}
	// TODO: Automatically rollback here if we failed above. Locate the last complete
	// revision and populate the rollback spec with it.
	// See https://github.com/kubernetes/kubernetes/issues/23211.

	if d.Spec.Paused {
		return dc.sync(d, rsList, podMap)
	}

	if d.Spec.RollbackTo != nil {
		revision := d.Spec.RollbackTo.Revision
		if d, err = dc.rollback(d, rsList, podMap, &revision); err != nil {
			return err
		}
	}

	scalingEvent, err := dc.isScalingEvent(d, rsList, podMap)
	if err != nil {
		return err
	}
	if scalingEvent {
		return dc.sync(d, rsList, podMap)
	}

	switch d.Spec.Strategy.Type {
	case extensions.RecreateDeploymentStrategyType:
		return dc.rolloutRecreate(d, rsList, podMap)
	case extensions.RollingUpdateDeploymentStrategyType:
		return dc.rolloutRolling(d, rsList, podMap)
	}
	return fmt.Errorf("unexpected deployment strategy type: %s", d.Spec.Strategy.Type)
}
