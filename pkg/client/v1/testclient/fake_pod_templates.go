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

package testclient

import (
	v1api "github.com/GoogleCloudPlatform/kubernetes/pkg/api/v1"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/fields"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/watch"
)

// FakePodTemplates implements PodTemplatesInterface. Meant to be embedded into a struct to get a default
// implementation. This makes faking out just the methods you want to test easier.
type FakePodTemplates struct {
	Fake      *Fake
	Namespace string
}

func (c *FakePodTemplates) List(label labels.Selector, field fields.Selector) (*v1api.PodTemplateList, error) {
	obj, err := c.Fake.Invokes(FakeAction{Action: "list-podTemplates"}, &v1api.PodTemplateList{})
	return obj.(*v1api.PodTemplateList), err
}

func (c *FakePodTemplates) Get(name string) (*v1api.PodTemplate, error) {
	obj, err := c.Fake.Invokes(FakeAction{Action: "get-podTemplate", Value: name}, &v1api.PodTemplate{})
	return obj.(*v1api.PodTemplate), err
}

func (c *FakePodTemplates) Delete(name string, options *v1api.DeleteOptions) error {
	_, err := c.Fake.Invokes(FakeAction{Action: "delete-podTemplate", Value: name}, &v1api.PodTemplate{})
	return err
}

func (c *FakePodTemplates) Create(pod *v1api.PodTemplate) (*v1api.PodTemplate, error) {
	obj, err := c.Fake.Invokes(FakeAction{Action: "create-podTemplate"}, &v1api.PodTemplate{})
	return obj.(*v1api.PodTemplate), err
}

func (c *FakePodTemplates) Update(pod *v1api.PodTemplate) (*v1api.PodTemplate, error) {
	obj, err := c.Fake.Invokes(FakeAction{Action: "update-podTemplate", Value: pod.Name}, &v1api.PodTemplate{})
	return obj.(*v1api.PodTemplate), err
}

func (c *FakePodTemplates) Watch(label labels.Selector, field fields.Selector, resourceVersion string) (watch.Interface, error) {
	c.Fake.Actions = append(c.Fake.Actions, FakeAction{Action: "watch-podTemplates", Value: resourceVersion})
	return c.Fake.Watch, c.Fake.Err
}
