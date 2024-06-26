/*
Copyright 2023 The bpfman Authors.

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
	"context"
	time "time"

	apisv1alpha1 "github.com/bpfman/bpfman-operator/apis/v1alpha1"
	v1alpha1 "github.com/bpfman/bpfman-operator/pkg/client/apis/v1alpha1"
	clientset "github.com/bpfman/bpfman-operator/pkg/client/clientset"
	internalinterfaces "github.com/bpfman/bpfman-operator/pkg/client/externalversions/internalinterfaces"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FexitProgramInformer provides access to a shared informer and lister for
// FexitPrograms.
type FexitProgramInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FexitProgramLister
}

type fexitProgramInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewFexitProgramInformer constructs a new informer for FexitProgram type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFexitProgramInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFexitProgramInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredFexitProgramInformer constructs a new informer for FexitProgram type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFexitProgramInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BpfmanV1alpha1().FexitPrograms().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BpfmanV1alpha1().FexitPrograms().Watch(context.TODO(), options)
			},
		},
		&apisv1alpha1.FexitProgram{},
		resyncPeriod,
		indexers,
	)
}

func (f *fexitProgramInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFexitProgramInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *fexitProgramInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisv1alpha1.FexitProgram{}, f.defaultInformer)
}

func (f *fexitProgramInformer) Lister() v1alpha1.FexitProgramLister {
	return v1alpha1.NewFexitProgramLister(f.Informer().GetIndexer())
}
