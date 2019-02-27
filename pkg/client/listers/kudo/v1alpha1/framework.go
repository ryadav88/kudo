/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kudobuilder/kudo/pkg/apis/kudo/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FrameworkLister helps list Frameworks.
type FrameworkLister interface {
	// List lists all Frameworks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Framework, err error)
	// Frameworks returns an object that can list and get Frameworks.
	Frameworks(namespace string) FrameworkNamespaceLister
	FrameworkListerExpansion
}

// frameworkLister implements the FrameworkLister interface.
type frameworkLister struct {
	indexer cache.Indexer
}

// NewFrameworkLister returns a new FrameworkLister.
func NewFrameworkLister(indexer cache.Indexer) FrameworkLister {
	return &frameworkLister{indexer: indexer}
}

// List lists all Frameworks in the indexer.
func (s *frameworkLister) List(selector labels.Selector) (ret []*v1alpha1.Framework, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Framework))
	})
	return ret, err
}

// Frameworks returns an object that can list and get Frameworks.
func (s *frameworkLister) Frameworks(namespace string) FrameworkNamespaceLister {
	return frameworkNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FrameworkNamespaceLister helps list and get Frameworks.
type FrameworkNamespaceLister interface {
	// List lists all Frameworks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Framework, err error)
	// Get retrieves the Framework from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Framework, error)
	FrameworkNamespaceListerExpansion
}

// frameworkNamespaceLister implements the FrameworkNamespaceLister
// interface.
type frameworkNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Frameworks in the indexer for a given namespace.
func (s frameworkNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Framework, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Framework))
	})
	return ret, err
}

// Get retrieves the Framework from the indexer for a given namespace and name.
func (s frameworkNamespaceLister) Get(name string) (*v1alpha1.Framework, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("framework"), name)
	}
	return obj.(*v1alpha1.Framework), nil
}
