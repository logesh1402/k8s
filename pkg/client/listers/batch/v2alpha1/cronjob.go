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

// This file was automatically generated by lister-gen

package v2alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	batch "k8s.io/apis/pkg/apis/batch"
	v2alpha1 "k8s.io/apis/pkg/apis/batch/v2alpha1"
	"k8s.io/client-go/tools/cache"
)

// CronJobLister helps list CronJobs.
type CronJobLister interface {
	// List lists all CronJobs in the indexer.
	List(selector labels.Selector) (ret []*v2alpha1.CronJob, err error)
	// CronJobs returns an object that can list and get CronJobs.
	CronJobs(namespace string) CronJobNamespaceLister
	CronJobListerExpansion
}

// cronJobLister implements the CronJobLister interface.
type cronJobLister struct {
	indexer cache.Indexer
}

// NewCronJobLister returns a new CronJobLister.
func NewCronJobLister(indexer cache.Indexer) CronJobLister {
	return &cronJobLister{indexer: indexer}
}

// List lists all CronJobs in the indexer.
func (s *cronJobLister) List(selector labels.Selector) (ret []*v2alpha1.CronJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2alpha1.CronJob))
	})
	return ret, err
}

// CronJobs returns an object that can list and get CronJobs.
func (s *cronJobLister) CronJobs(namespace string) CronJobNamespaceLister {
	return cronJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CronJobNamespaceLister helps list and get CronJobs.
type CronJobNamespaceLister interface {
	// List lists all CronJobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v2alpha1.CronJob, err error)
	// Get retrieves the CronJob from the indexer for a given namespace and name.
	Get(name string) (*v2alpha1.CronJob, error)
	CronJobNamespaceListerExpansion
}

// cronJobNamespaceLister implements the CronJobNamespaceLister
// interface.
type cronJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CronJobs in the indexer for a given namespace.
func (s cronJobNamespaceLister) List(selector labels.Selector) (ret []*v2alpha1.CronJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2alpha1.CronJob))
	})
	return ret, err
}

// Get retrieves the CronJob from the indexer for a given namespace and name.
func (s cronJobNamespaceLister) Get(name string) (*v2alpha1.CronJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(batch.Resource("cronjob"), name)
	}
	return obj.(*v2alpha1.CronJob), nil
}
