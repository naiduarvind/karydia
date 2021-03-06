// Copyright (C) 2019 SAP SE or an SAP affiliate company. All rights reserved.
// This file is licensed under the Apache Software License, v. 2 except as
// noted otherwise in the LICENSE file.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/karydia/karydia/pkg/apis/karydia/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KarydiaConfigLister helps list KarydiaConfigs.
type KarydiaConfigLister interface {
	// List lists all KarydiaConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.KarydiaConfig, err error)
	// Get retrieves the KarydiaConfig from the index for a given name.
	Get(name string) (*v1alpha1.KarydiaConfig, error)
	KarydiaConfigListerExpansion
}

// karydiaConfigLister implements the KarydiaConfigLister interface.
type karydiaConfigLister struct {
	indexer cache.Indexer
}

// NewKarydiaConfigLister returns a new KarydiaConfigLister.
func NewKarydiaConfigLister(indexer cache.Indexer) KarydiaConfigLister {
	return &karydiaConfigLister{indexer: indexer}
}

// List lists all KarydiaConfigs in the indexer.
func (s *karydiaConfigLister) List(selector labels.Selector) (ret []*v1alpha1.KarydiaConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KarydiaConfig))
	})
	return ret, err
}

// Get retrieves the KarydiaConfig from the index for a given name.
func (s *karydiaConfigLister) Get(name string) (*v1alpha1.KarydiaConfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("karydiaconfig"), name)
	}
	return obj.(*v1alpha1.KarydiaConfig), nil
}
