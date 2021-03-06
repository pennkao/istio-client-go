/*
Portions Copyright 2019 The Kubernetes Authors.
Portions Copyright 2019 Aspen Mesh Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	rbacv1alpha1 "github.com/aspenmesh/istio-client-go/pkg/apis/rbac/v1alpha1"
	versioned "github.com/aspenmesh/istio-client-go/pkg/client/clientset/versioned"
	internalinterfaces "github.com/aspenmesh/istio-client-go/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/aspenmesh/istio-client-go/pkg/client/listers/rbac/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceRoleBindingInformer provides access to a shared informer and lister for
// ServiceRoleBindings.
type ServiceRoleBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ServiceRoleBindingLister
}

type serviceRoleBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServiceRoleBindingInformer constructs a new informer for ServiceRoleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceRoleBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceRoleBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServiceRoleBindingInformer constructs a new informer for ServiceRoleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceRoleBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RbacV1alpha1().ServiceRoleBindings(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RbacV1alpha1().ServiceRoleBindings(namespace).Watch(options)
			},
		},
		&rbacv1alpha1.ServiceRoleBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceRoleBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceRoleBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceRoleBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rbacv1alpha1.ServiceRoleBinding{}, f.defaultInformer)
}

func (f *serviceRoleBindingInformer) Lister() v1alpha1.ServiceRoleBindingLister {
	return v1alpha1.NewServiceRoleBindingLister(f.Informer().GetIndexer())
}
