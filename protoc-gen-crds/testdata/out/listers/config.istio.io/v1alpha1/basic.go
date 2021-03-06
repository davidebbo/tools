// Copyright 2018 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "istio.io/tools/protoc-gen-crds/testdata/out/config.istio.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BasicLister helps list Basics.
type BasicLister interface {
	// List lists all Basics in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Basic, err error)
	// Basics returns an object that can list and get Basics.
	Basics(namespace string) BasicNamespaceLister
	BasicListerExpansion
}

// basicLister implements the BasicLister interface.
type basicLister struct {
	indexer cache.Indexer
}

// NewBasicLister returns a new BasicLister.
func NewBasicLister(indexer cache.Indexer) BasicLister {
	return &basicLister{indexer: indexer}
}

// List lists all Basics in the indexer.
func (s *basicLister) List(selector labels.Selector) (ret []*v1alpha1.Basic, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Basic))
	})
	return ret, err
}

// Basics returns an object that can list and get Basics.
func (s *basicLister) Basics(namespace string) BasicNamespaceLister {
	return basicNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BasicNamespaceLister helps list and get Basics.
type BasicNamespaceLister interface {
	// List lists all Basics in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Basic, err error)
	// Get retrieves the Basic from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Basic, error)
	BasicNamespaceListerExpansion
}

// basicNamespaceLister implements the BasicNamespaceLister
// interface.
type basicNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Basics in the indexer for a given namespace.
func (s basicNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Basic, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Basic))
	})
	return ret, err
}

// Get retrieves the Basic from the indexer for a given namespace and name.
func (s basicNamespaceLister) Get(name string) (*v1alpha1.Basic, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("basic"), name)
	}
	return obj.(*v1alpha1.Basic), nil
}
