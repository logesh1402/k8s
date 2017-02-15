/*
Copyright 2017 The Kubernetes Authors.

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

package upgrades

import (
	"fmt"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/kubernetes/pkg/controller"

	"k8s.io/kubernetes/pkg/api/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	extensions "k8s.io/kubernetes/pkg/apis/extensions/v1beta1"
	"k8s.io/kubernetes/test/e2e/framework"

	. "github.com/onsi/ginkgo"
)

// DaemonSetUpgradeTest tests that a DaemonSet is running before, during and after
// a cluster upgrade.
type DaemonSetUpgradeTest struct {
	daemonSet *extensions.DaemonSet
}

// Setup creates a DaemonSet and verifies that it's running
func (t *DaemonSetUpgradeTest) Setup(f *framework.Framework) {
	namespace := "daemonset-upgrade"
	daemonSetName := "ds1"
	labelSet := map[string]string{"ds-name": daemonSetName}
	image := "gcr.io/google_containers/serve_hostname:v1.4"

	// Grab a unique namespace so we don't collide.
	ns, err := f.CreateNamespace(namespace, nil)
	framework.ExpectNoError(err)

	t.daemonSet = &extensions.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns.Name,
			Name:      daemonSetName,
		},
		Spec: extensions.DaemonSetSpec{
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labelSet,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  daemonSetName,
							Image: image,
							Ports: []v1.ContainerPort{{ContainerPort: 9376}},
						},
					},
				},
			},
		},
	}

	By("Creating a DaemonSet")
	if t.daemonSet, err = f.ClientSet.Extensions().DaemonSets(ns.Name).Create(t.daemonSet); err != nil {
		framework.Failf("unable to create test DaemonSet %s: %v", t.daemonSet.Name, err)
	}

	By("Waiting for DaemonSet pods to become ready")
	wait.Poll(framework.Poll, framework.PodStartTimeout, func() (bool, error) {
		return checkRunningOnAllNodes(f, t.daemonSet.Namespace, t.daemonSet.Labels)
	})
	framework.ExpectNoError(err)

	By("Validating the DaemonSet after creation")
	t.validateRunningDaemonSet(f)
}

// Test validates that the DaemonSet is running during the upgrade (if applicable) and
// verifies again post-upgrade that it's running
func (t *DaemonSetUpgradeTest) Test(f *framework.Framework, done <-chan struct{}, upgrade UpgradeType) {
	testDuringDisruption := upgrade == MasterUpgrade

	if testDuringDisruption {
		By("validating the DaemonSet is still running during upgrade")

		wait.Until(func() {
			t.validateRunningDaemonSet(f)
		}, framework.Poll, done)
	}

	<-done

	By("validating the DaemonSet is still running after upgrade")
	t.validateRunningDaemonSet(f)
}

// Teardown cleans up any remaining resources.
func (t *DaemonSetUpgradeTest) Teardown(f *framework.Framework) {
	// rely on the namespace deletion to clean up everything
}

func (t *DaemonSetUpgradeTest) validateRunningDaemonSet(f *framework.Framework) {
	// Pods should come up in a reasonable amount of time
	By("confirming the DaemonSet pods are running on all expected nodes")
	_, err := checkRunningOnAllNodes(f, t.daemonSet.Namespace, t.daemonSet.Labels)
	framework.ExpectNoError(err)

	// DaemonSet resource itself should be good
	By("confirming the DaemonSet resource is in a good state")
	err = checkDaemonStatus(f, t.daemonSet.Namespace, t.daemonSet.Name)
	framework.ExpectNoError(err)
}

func checkRunningOnAllNodes(f *framework.Framework, namespace string, selector map[string]string) (bool, error) {

	nodeList, err := f.ClientSet.Core().Nodes().List(metav1.ListOptions{})
	framework.ExpectNoError(err)
	nodeNames := make([]string, 0)
	for _, node := range nodeList.Items {
		taints, err := v1.GetNodeTaints(&node)
		if err == nil && len(taints) == 0 {
			nodeNames = append(nodeNames, node.Name)
		} else {
			framework.Logf("Node %v not expected to have DaemonSet pod, has taints %v", node.Name, taints)
		}
	}

	return checkDaemonPodOnNodes(f, namespace, selector, nodeNames)
}

func checkDaemonPodOnNodes(f *framework.Framework, namespace string, labelSet map[string]string, nodeNames []string) (bool, error) {
	selector := labels.Set(labelSet).AsSelector()
	options := metav1.ListOptions{LabelSelector: selector.String()}
	podList, err := f.ClientSet.Core().Pods(namespace).List(options)
	if err != nil {
		return false, nil
	}
	pods := podList.Items

	nodesToPodCount := make(map[string]int)
	for _, pod := range pods {
		if controller.IsPodActive(&pod) {
			framework.Logf("Pod name: %v\t Node Name: %v", pod.Name, pod.Spec.NodeName)
			nodesToPodCount[pod.Spec.NodeName]++
		}
	}
	framework.Logf("nodesToPodCount: %v", nodesToPodCount)

	// Ensure that exactly 1 pod is running on all nodes in nodeNames.
	for _, nodeName := range nodeNames {
		if nodesToPodCount[nodeName] != 1 {
			return false, nil
		}
	}

	// Ensure that sizes of the lists are the same. We've verified that every element of nodeNames is in
	// nodesToPodCount, so verifying the lengths are equal ensures that there aren't pods running on any
	// other nodes.
	return len(nodesToPodCount) == len(nodeNames), nil
}

func checkDaemonStatus(f *framework.Framework, namespace string, dsName string) error {
	ds, err := f.ClientSet.ExtensionsV1beta1().DaemonSets(namespace).Get(dsName, metav1.GetOptions{})
	framework.ExpectNoError(err)

	desired, scheduled, ready := ds.Status.DesiredNumberScheduled, ds.Status.CurrentNumberScheduled, ds.Status.NumberReady
	if desired != scheduled && desired != ready {
		return fmt.Errorf("Error in daemon status. DesiredScheduled: %d, CurrentScheduled: %d, Ready: %d", desired, scheduled, ready)
	}

	return nil
}
