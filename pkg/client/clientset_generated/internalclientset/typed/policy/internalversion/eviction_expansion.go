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

package internalversion

import (
	"strings"

	policy "k8s.io/kubernetes/pkg/apis/policy"
	"k8s.io/kubernetes/pkg/client/restclient"
)

const (
	PolicyAPIVersion    = "policy/v1beta1"
	EvictionKind        = "Eviction"
	EvictionSubresource = "pods/eviction"
)

// The EvictionExpansion interface allows manually adding extra methods to the ScaleInterface.
type EvictionExpansion interface {
	Evict(eviction *policy.Eviction) error
}

func (c *evictions) Evict(eviction *policy.Eviction) error {
	client := c.client.(*restclient.RESTClient)
	originalVersionedAPIPath := client.VersionedAPIPath
	client.VersionedAPIPath = strings.Replace(originalVersionedAPIPath, "/apis/"+PolicyAPIVersion, "/api/v1", 1)
	err := client.Post().
		Namespace(eviction.Namespace).
		Resource("pods").
		Name(eviction.Name).
		SubResource("eviction").
		Body(eviction).
		Do().
		Error()
	client.VersionedAPIPath = originalVersionedAPIPath
	return err
}
