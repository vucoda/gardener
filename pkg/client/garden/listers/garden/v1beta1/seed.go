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

package v1beta1

import (
	v1beta1 "github.com/gardener/gardener/pkg/apis/garden/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SeedLister helps list Seeds.
type SeedLister interface {
	// List lists all Seeds in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.Seed, err error)
	// Get retrieves the Seed from the index for a given name.
	Get(name string) (*v1beta1.Seed, error)
	SeedListerExpansion
}

// seedLister implements the SeedLister interface.
type seedLister struct {
	indexer cache.Indexer
}

// NewSeedLister returns a new SeedLister.
func NewSeedLister(indexer cache.Indexer) SeedLister {
	return &seedLister{indexer: indexer}
}

// List lists all Seeds in the indexer.
func (s *seedLister) List(selector labels.Selector) (ret []*v1beta1.Seed, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Seed))
	})
	return ret, err
}

// Get retrieves the Seed from the index for a given name.
func (s *seedLister) Get(name string) (*v1beta1.Seed, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("seed"), name)
	}
	return obj.(*v1beta1.Seed), nil
}
