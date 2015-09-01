/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package etcd

import (
	"testing"
	"time"

	"github.com/coreos/go-etcd/etcd"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/rest/resttest"
	"k8s.io/kubernetes/pkg/api/testapi"
	"k8s.io/kubernetes/pkg/expapi"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	etcdgeneric "k8s.io/kubernetes/pkg/registry/generic/etcd"
	"k8s.io/kubernetes/pkg/registry/registrytest"
	"k8s.io/kubernetes/pkg/runtime"
	etcdstorage "k8s.io/kubernetes/pkg/storage/etcd"
	"k8s.io/kubernetes/pkg/tools"
	"k8s.io/kubernetes/pkg/tools/etcdtest"
)

const (
	PASS = iota
	FAIL
)

func newStorage(t *testing.T) (*REST, *tools.FakeEtcdClient) {
	etcdStorage, fakeClient := registrytest.NewEtcdStorage(t)
	return NewREST(etcdStorage), fakeClient
}

func validNewDeployment() *expapi.Deployment {
	return &expapi.Deployment{
		ObjectMeta: api.ObjectMeta{
			Name:      "foo",
			Namespace: api.NamespaceDefault,
		},
		Spec: expapi.DeploymentSpec{
			Selector: map[string]string{"a": "b"},
			Template: &api.PodTemplateSpec{
				ObjectMeta: api.ObjectMeta{
					Labels: map[string]string{"a": "b"},
				},
				Spec: api.PodSpec{
					Containers: []api.Container{
						{
							Name:            "test",
							Image:           "test_image",
							ImagePullPolicy: api.PullIfNotPresent,
						},
					},
					RestartPolicy: api.RestartPolicyAlways,
					DNSPolicy:     api.DNSClusterFirst,
				},
			},
			UniqueLabelKey: "my-label",
		},
	}
}

var validDeployment = *validNewDeployment()

func TestCreate(t *testing.T) {
	storage, fakeClient := newStorage(t)
	test := resttest.New(t, storage, fakeClient.SetError)
	deployment := validNewDeployment()
	deployment.ObjectMeta = api.ObjectMeta{}
	test.TestCreate(
		// valid
		deployment,
		func(ctx api.Context, obj runtime.Object) error {
			return registrytest.SetObject(fakeClient, storage.KeyFunc, ctx, obj)
		},
		func(ctx api.Context, obj runtime.Object) (runtime.Object, error) {
			return registrytest.GetObject(fakeClient, storage.KeyFunc, storage.NewFunc, ctx, obj)
		},
		// invalid (invalid selector)
		&expapi.Deployment{
			Spec: expapi.DeploymentSpec{
				Selector: map[string]string{},
				Template: validDeployment.Spec.Template,
			},
		},
	)
}

func TestUpdate(t *testing.T) {
	storage, fakeClient := newStorage(t)
	test := resttest.New(t, storage, fakeClient.SetError)
	test.TestUpdate(
		// valid
		validNewDeployment(),
		func(ctx api.Context, obj runtime.Object) error {
			return registrytest.SetObject(fakeClient, storage.KeyFunc, ctx, obj)
		},
		func(resourceVersion uint64) {
			registrytest.SetResourceVersion(fakeClient, resourceVersion)
		},
		func(ctx api.Context, obj runtime.Object) (runtime.Object, error) {
			return registrytest.GetObject(fakeClient, storage.KeyFunc, storage.NewFunc, ctx, obj)
		},
		// updateFunc
		func(obj runtime.Object) runtime.Object {
			object := obj.(*expapi.Deployment)
			object.Spec.Template.Spec.NodeSelector = map[string]string{"c": "d"}
			return object
		},
		// invalid updateFunc
		func(obj runtime.Object) runtime.Object {
			object := obj.(*expapi.Deployment)
			object.UID = "newUID"
			return object
		},
		func(obj runtime.Object) runtime.Object {
			object := obj.(*expapi.Deployment)
			object.Name = ""
			return object
		},
		func(obj runtime.Object) runtime.Object {
			object := obj.(*expapi.Deployment)
			object.Spec.Template.Spec.RestartPolicy = api.RestartPolicyOnFailure
			return object
		},
		func(obj runtime.Object) runtime.Object {
			object := obj.(*expapi.Deployment)
			object.Spec.Selector = map[string]string{}
			return object
		},
	)
}

func TestEtcdGet(t *testing.T) {
	storage, fakeClient := newStorage(t)
	test := resttest.New(t, storage, fakeClient.SetError)
	test.TestGet(validNewDeployment())
}

func TestEtcdList(t *testing.T) {
	storage, fakeClient := newStorage(t)
	test := resttest.New(t, storage, fakeClient.SetError)
	key := etcdtest.AddPrefix(storage.KeyRootFunc(test.TestContext()))
	test.TestList(
		validNewDeployment(),
		func(objects []runtime.Object) []runtime.Object {
			return registrytest.SetObjectsForKey(fakeClient, key, objects)
		},
		func(resourceVersion uint64) {
			registrytest.SetResourceVersion(fakeClient, resourceVersion)
		})
}

func TestEtcdDelete(t *testing.T) {
	ctx := api.NewDefaultContext()
	storage, fakeClient := newStorage(t)
	key, err := storage.KeyFunc(ctx, validDeployment.Name)
	key = etcdtest.AddPrefix(key)

	fakeClient.Set(key, runtime.EncodeOrDie(testapi.Codec(), validNewDeployment()), 0)
	obj, err := storage.Delete(ctx, validDeployment.Name, nil)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if status, ok := obj.(*api.Status); !ok {
		t.Errorf("Expected status of delete, got %#v", status)
	} else if status.Status != api.StatusSuccess {
		t.Errorf("Expected success, got %#v", status.Status)
	}
	if len(fakeClient.DeletedKeys) != 1 {
		t.Errorf("Expected 1 delete, found %#v", fakeClient.DeletedKeys)
	}
	if fakeClient.DeletedKeys[0] != key {
		t.Errorf("Unexpected key: %s, expected %s", fakeClient.DeletedKeys[0], key)
	}
}

func TestEtcdWatch(t *testing.T) {
	ctx := api.NewDefaultContext()
	storage, fakeClient := newStorage(t)
	watching, err := storage.Watch(ctx,
		labels.Everything(),
		fields.Everything(),
		"1",
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	fakeClient.WaitForWatchCompletion()

	select {
	case _, ok := <-watching.ResultChan():
		if !ok {
			t.Errorf("watching channel should be open")
		}
	default:
	}
	fakeClient.WatchInjectError <- nil
	if _, ok := <-watching.ResultChan(); ok {
		t.Errorf("watching channel should be closed")
	}
	watching.Stop()
}

// Tests that we can watch for the creation of Deployment with specified labels.
func TestEtcdWatchWithLabels(t *testing.T) {
	ctx := api.WithNamespace(api.NewDefaultContext(), validDeployment.Namespace)
	storage, fakeClient := newStorage(t)
	fakeClient.ExpectNotFoundGet(etcdgeneric.NamespaceKeyRootFunc(ctx, "/registry/pods"))

	watching, err := storage.Watch(ctx,
		labels.SelectorFromSet(validDeployment.Spec.Selector),
		fields.Everything(),
		"1",
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	fakeClient.WaitForWatchCompletion()

	// The watcher above is waiting for these Labels, on receiving them it should
	// apply the deploymentStatus decorator, which lists pods, causing a query against
	// the /registry/pods endpoint of the etcd client.
	deployment := &expapi.Deployment{
		ObjectMeta: api.ObjectMeta{
			Name:      "foo",
			Labels:    validDeployment.Spec.Selector,
			Namespace: "default",
		},
	}
	deploymentBytes, _ := testapi.Codec().Encode(deployment)
	fakeClient.WatchResponse <- &etcd.Response{
		Action: "create",
		Node: &etcd.Node{
			Value: string(deploymentBytes),
		},
	}
	select {
	case _, ok := <-watching.ResultChan():
		if !ok {
			t.Errorf("watching channel should be open")
		}
	case <-time.After(time.Millisecond * 100):
		t.Error("unexpected timeout from result channel")
	}
	watching.Stop()
}

// Tests that we can watch for Deployment with specified fields.
func TestEtcdWatchWithFields(t *testing.T) {
	ctx := api.WithNamespace(api.NewDefaultContext(), validDeployment.Namespace)
	storage, fakeClient := newStorage(t)
	fakeClient.ExpectNotFoundGet(etcdgeneric.NamespaceKeyRootFunc(ctx, "/registry/pods"))

	testFieldMap := map[int][]fields.Set{
		PASS: {
			{"metadata.name": "foo"},
		},
		FAIL: {
			{"metadata.name": "bar"},
			{"name": "foo"},
		},
	}
	testEtcdActions := []string{
		etcdstorage.EtcdCreate,
		etcdstorage.EtcdSet,
		etcdstorage.EtcdCAS,
		etcdstorage.EtcdDelete}

	deployment := &expapi.Deployment{
		ObjectMeta: api.ObjectMeta{
			Name:      "foo",
			Labels:    validDeployment.Spec.Selector,
			Namespace: "default",
		},
		Status: expapi.DeploymentStatus{
			Replicas:        1,
			UpdatedReplicas: 4,
		},
	}
	deploymentBytes, _ := testapi.Codec().Encode(deployment)

	for expectedResult, fieldSet := range testFieldMap {
		for _, field := range fieldSet {
			for _, action := range testEtcdActions {
				watching, err := storage.Watch(ctx,
					labels.Everything(),
					field.AsSelector(),
					"1",
				)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				var prevNode *etcd.Node = nil
				node := &etcd.Node{
					Value: string(deploymentBytes),
				}
				if action == etcdstorage.EtcdDelete {
					prevNode = node
				}
				fakeClient.WaitForWatchCompletion()
				fakeClient.WatchResponse <- &etcd.Response{
					Action:   action,
					Node:     node,
					PrevNode: prevNode,
				}

				select {
				case r, ok := <-watching.ResultChan():
					if expectedResult == FAIL {
						t.Errorf("Unexpected result from channel %#v. Field: %v", r, field)
					}
					if !ok {
						t.Errorf("watching channel should be open")
					}
				case <-time.After(time.Millisecond * 100):
					if expectedResult == PASS {
						t.Error("unexpected timeout from result channel")
					}
				}
				watching.Stop()
			}
		}
	}
}

func TestEtcdWatchNotMatch(t *testing.T) {
	ctx := api.NewDefaultContext()
	storage, fakeClient := newStorage(t)
	fakeClient.ExpectNotFoundGet(etcdgeneric.NamespaceKeyRootFunc(ctx, "/registry/pods"))

	watching, err := storage.Watch(ctx,
		labels.SelectorFromSet(labels.Set{"name": "foo"}),
		fields.Everything(),
		"1",
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	fakeClient.WaitForWatchCompletion()

	deployment := &expapi.Deployment{
		ObjectMeta: api.ObjectMeta{
			Name: "bar",
			Labels: map[string]string{
				"name": "bar",
			},
		},
	}
	deploymentBytes, _ := testapi.Codec().Encode(deployment)
	fakeClient.WatchResponse <- &etcd.Response{
		Action: "create",
		Node: &etcd.Node{
			Value: string(deploymentBytes),
		},
	}

	select {
	case <-watching.ResultChan():
		t.Error("unexpected result from result channel")
	case <-time.After(time.Millisecond * 100):
		// expected case
	}
}

func TestDelete(t *testing.T) {
	ctx := api.NewDefaultContext()
	storage, fakeClient := newStorage(t)
	test := resttest.New(t, storage, fakeClient.SetError)
	key, _ := storage.KeyFunc(ctx, validDeployment.Name)
	key = etcdtest.AddPrefix(key)

	createFn := func() runtime.Object {
		dc := validNewDeployment()
		dc.ResourceVersion = "1"
		fakeClient.Data[key] = tools.EtcdResponseWithError{
			R: &etcd.Response{
				Node: &etcd.Node{
					Value:         runtime.EncodeOrDie(testapi.Codec(), dc),
					ModifiedIndex: 1,
				},
			},
		}
		return dc
	}
	gracefulSetFn := func() bool {
		// If the deployment is still around after trying to delete either the delete
		// failed, or we're deleting it gracefully.
		if fakeClient.Data[key].R.Node != nil {
			return true
		}
		return false
	}
	test.TestDelete(createFn, gracefulSetFn)
}
