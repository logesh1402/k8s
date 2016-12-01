/*
Copyright 2016 The Kubernetes Authors.

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
	"k8s.io/kubernetes/pkg/apis/rbac"
	"k8s.io/kubernetes/pkg/genericapiserver"
	"k8s.io/kubernetes/pkg/registry/generic/registry"
	"k8s.io/kubernetes/pkg/registry/rbac/rolebinding"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/storage"
	"k8s.io/kubernetes/pkg/util/restoptions"
)

// REST implements a RESTStorage for RoleBinding against etcd
type REST struct {
	*registry.Store
}

// NewREST returns a RESTStorage object that will work against RoleBinding objects.
func NewREST(optsGetter genericapiserver.RESTOptionsGetter) *REST {
	store := &registry.Store{
		NewFunc:     func() runtime.Object { return &rbac.RoleBinding{} },
		NewListFunc: func() runtime.Object { return &rbac.RoleBindingList{} },
		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*rbac.RoleBinding).Name, nil
		},
		PredicateFunc:     rolebinding.Matcher,
		QualifiedResource: rbac.Resource("rolebindings"),

		CreateStrategy: rolebinding.Strategy,
		UpdateStrategy: rolebinding.Strategy,
		DeleteStrategy: rolebinding.Strategy,
	}
	restoptions.ApplyOptions(optsGetter, store, storage.NoTriggerPublisher)

	return &REST{store}
}
