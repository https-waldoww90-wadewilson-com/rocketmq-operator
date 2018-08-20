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

package v1alpha1

import (
	time "time"

	rocketmqv1alpha1 "github.com/huanwei/rocketmq-operator/pkg/apis/rocketmq/v1alpha1"
	versioned "github.com/huanwei/rocketmq-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/huanwei/rocketmq-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/huanwei/rocketmq-operator/pkg/generated/listers/rocketmq/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BrokerClusterInformer provides access to a shared informer and lister for
// BrokerClusters.
type BrokerClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BrokerClusterLister
}

type brokerClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBrokerClusterInformer constructs a new informer for BrokerCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBrokerClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBrokerClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBrokerClusterInformer constructs a new informer for BrokerCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBrokerClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RocketmqV1alpha1().BrokerClusters(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RocketmqV1alpha1().BrokerClusters(namespace).Watch(options)
			},
		},
		&rocketmqv1alpha1.BrokerCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *brokerClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBrokerClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *brokerClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rocketmqv1alpha1.BrokerCluster{}, f.defaultInformer)
}

func (f *brokerClusterInformer) Lister() v1alpha1.BrokerClusterLister {
	return v1alpha1.NewBrokerClusterLister(f.Informer().GetIndexer())
}
