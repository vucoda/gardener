/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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
	v1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ControlPlaneLister helps list ControlPlanes.
type ControlPlaneLister interface {
	// List lists all ControlPlanes in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ControlPlane, err error)
	// ControlPlanes returns an object that can list and get ControlPlanes.
	ControlPlanes(namespace string) ControlPlaneNamespaceLister
	ControlPlaneListerExpansion
}

// controlPlaneLister implements the ControlPlaneLister interface.
type controlPlaneLister struct {
	indexer cache.Indexer
}

// NewControlPlaneLister returns a new ControlPlaneLister.
func NewControlPlaneLister(indexer cache.Indexer) ControlPlaneLister {
	return &controlPlaneLister{indexer: indexer}
}

// List lists all ControlPlanes in the indexer.
func (s *controlPlaneLister) List(selector labels.Selector) (ret []*v1alpha1.ControlPlane, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ControlPlane))
	})
	return ret, err
}

// ControlPlanes returns an object that can list and get ControlPlanes.
func (s *controlPlaneLister) ControlPlanes(namespace string) ControlPlaneNamespaceLister {
	return controlPlaneNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ControlPlaneNamespaceLister helps list and get ControlPlanes.
type ControlPlaneNamespaceLister interface {
	// List lists all ControlPlanes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ControlPlane, err error)
	// Get retrieves the ControlPlane from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ControlPlane, error)
	ControlPlaneNamespaceListerExpansion
}

// controlPlaneNamespaceLister implements the ControlPlaneNamespaceLister
// interface.
type controlPlaneNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ControlPlanes in the indexer for a given namespace.
func (s controlPlaneNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ControlPlane, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ControlPlane))
	})
	return ret, err
}

// Get retrieves the ControlPlane from the indexer for a given namespace and name.
func (s controlPlaneNamespaceLister) Get(name string) (*v1alpha1.ControlPlane, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("controlplane"), name)
	}
	return obj.(*v1alpha1.ControlPlane), nil
}
