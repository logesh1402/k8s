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

package client

import (
	"net/url"
	"testing"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/api/resource"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api/testapi"
	v1api "github.com/GoogleCloudPlatform/kubernetes/pkg/api/v1"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/fields"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
)

func getNodesResourceName() string {
	return "nodes"
}

func TestListMinions(t *testing.T) {
	c := &testClient{
		Request: testRequest{
			Method: "GET",
			Path:   testapi.ResourcePath(getNodesResourceName(), "", ""),
		},
		Response: Response{StatusCode: 200, Body: &v1api.NodeList{ListMeta: v1api.ListMeta{ResourceVersion: "1"}}},
	}
	response, err := c.Setup().Nodes().List(labels.Everything(), fields.Everything())
	c.Validate(t, response, err)
}

func TestListMinionsLabels(t *testing.T) {
	ns := "" // Value of api.NamespaceNone at the time v1 was made
	labelSelectorQueryParamName := "labelSelector"
	c := &testClient{
		Request: testRequest{
			Method: "GET",
			Path:   testapi.ResourcePath(getNodesResourceName(), "", ""),
			Query:  buildQueryValues(ns, url.Values{labelSelectorQueryParamName: []string{"foo=bar,name=baz"}})},
		Response: Response{
			StatusCode: 200,
			Body: &v1api.NodeList{
				Items: []v1api.Node{
					{
						ObjectMeta: v1api.ObjectMeta{
							Labels: map[string]string{
								"foo":  "bar",
								"name": "baz",
							},
						},
					},
				},
			},
		},
	}
	c.Setup()
	c.QueryValidator[labelSelectorQueryParamName] = validateLabels
	selector := labels.Set{"foo": "bar", "name": "baz"}.AsSelector()
	receivedNodeList, err := c.Nodes().List(selector, fields.Everything())
	c.Validate(t, receivedNodeList, err)
}

func TestGetMinion(t *testing.T) {
	c := &testClient{
		Request: testRequest{
			Method: "GET",
			Path:   testapi.ResourcePath(getNodesResourceName(), "", "1"),
		},
		Response: Response{StatusCode: 200, Body: &v1api.Node{ObjectMeta: v1api.ObjectMeta{Name: "minion-1"}}},
	}
	response, err := c.Setup().Nodes().Get("1")
	c.Validate(t, response, err)
}

func TestGetMinionWithNoName(t *testing.T) {
	c := &testClient{Error: true}
	receivedNode, err := c.Setup().Nodes().Get("")
	if (err != nil) && (err.Error() != nameRequiredError) {
		t.Errorf("Expected error: %v, but got %v", nameRequiredError, err)
	}

	c.Validate(t, receivedNode, err)
}

func TestCreateMinion(t *testing.T) {
	requestMinion := &v1api.Node{
		ObjectMeta: v1api.ObjectMeta{
			Name: "minion-1",
		},
		Status: v1api.NodeStatus{
			Capacity: v1api.ResourceList{
				v1api.ResourceCPU:    resource.MustParse("1000m"),
				v1api.ResourceMemory: resource.MustParse("1Mi"),
			},
		},
		Spec: v1api.NodeSpec{
			Unschedulable: false,
		},
	}
	c := &testClient{
		Request: testRequest{
			Method: "POST",
			Path:   testapi.ResourcePath(getNodesResourceName(), "", ""),
			Body:   requestMinion},
		Response: Response{
			StatusCode: 200,
			Body:       requestMinion,
		},
	}
	receivedMinion, err := c.Setup().Nodes().Create(requestMinion)
	c.Validate(t, receivedMinion, err)
}

func TestDeleteMinion(t *testing.T) {
	c := &testClient{
		Request: testRequest{
			Method: "DELETE",
			Path:   testapi.ResourcePath(getNodesResourceName(), "", "foo"),
		},
		Response: Response{StatusCode: 200},
	}
	err := c.Setup().Nodes().Delete("foo")
	c.Validate(t, nil, err)
}

func TestUpdateMinion(t *testing.T) {
	requestMinion := &v1api.Node{
		ObjectMeta: v1api.ObjectMeta{
			Name:            "foo",
			ResourceVersion: "1",
		},
		Status: v1api.NodeStatus{
			Capacity: v1api.ResourceList{
				v1api.ResourceCPU:    resource.MustParse("1000m"),
				v1api.ResourceMemory: resource.MustParse("1Mi"),
			},
		},
		Spec: v1api.NodeSpec{
			Unschedulable: true,
		},
	}
	c := &testClient{
		Request: testRequest{
			Method: "PUT",
			Path:   testapi.ResourcePath(getNodesResourceName(), "", "foo"),
		},
		Response: Response{StatusCode: 200, Body: requestMinion},
	}
	response, err := c.Setup().Nodes().Update(requestMinion)
	c.Validate(t, response, err)
}
