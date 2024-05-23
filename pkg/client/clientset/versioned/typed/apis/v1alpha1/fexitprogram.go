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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/bpfman/bpfman-operator/apis/v1alpha1"
	scheme "github.com/bpfman/bpfman-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FexitProgramsGetter has a method to return a FexitProgramInterface.
// A group's client should implement this interface.
type FexitProgramsGetter interface {
	FexitPrograms() FexitProgramInterface
}

// FexitProgramInterface has methods to work with FexitProgram resources.
type FexitProgramInterface interface {
	Create(ctx context.Context, fexitProgram *v1alpha1.FexitProgram, opts v1.CreateOptions) (*v1alpha1.FexitProgram, error)
	Update(ctx context.Context, fexitProgram *v1alpha1.FexitProgram, opts v1.UpdateOptions) (*v1alpha1.FexitProgram, error)
	UpdateStatus(ctx context.Context, fexitProgram *v1alpha1.FexitProgram, opts v1.UpdateOptions) (*v1alpha1.FexitProgram, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.FexitProgram, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.FexitProgramList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FexitProgram, err error)
	FexitProgramExpansion
}

// fexitPrograms implements FexitProgramInterface
type fexitPrograms struct {
	client rest.Interface
}

// newFexitPrograms returns a FexitPrograms
func newFexitPrograms(c *BpfmanV1alpha1Client) *fexitPrograms {
	return &fexitPrograms{
		client: c.RESTClient(),
	}
}

// Get takes name of the fexitProgram, and returns the corresponding fexitProgram object, and an error if there is any.
func (c *fexitPrograms) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FexitProgram, err error) {
	result = &v1alpha1.FexitProgram{}
	err = c.client.Get().
		Resource("fexitprograms").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FexitPrograms that match those selectors.
func (c *fexitPrograms) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FexitProgramList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FexitProgramList{}
	err = c.client.Get().
		Resource("fexitprograms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested fexitPrograms.
func (c *fexitPrograms) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("fexitprograms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a fexitProgram and creates it.  Returns the server's representation of the fexitProgram, and an error, if there is any.
func (c *fexitPrograms) Create(ctx context.Context, fexitProgram *v1alpha1.FexitProgram, opts v1.CreateOptions) (result *v1alpha1.FexitProgram, err error) {
	result = &v1alpha1.FexitProgram{}
	err = c.client.Post().
		Resource("fexitprograms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fexitProgram).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a fexitProgram and updates it. Returns the server's representation of the fexitProgram, and an error, if there is any.
func (c *fexitPrograms) Update(ctx context.Context, fexitProgram *v1alpha1.FexitProgram, opts v1.UpdateOptions) (result *v1alpha1.FexitProgram, err error) {
	result = &v1alpha1.FexitProgram{}
	err = c.client.Put().
		Resource("fexitprograms").
		Name(fexitProgram.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fexitProgram).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *fexitPrograms) UpdateStatus(ctx context.Context, fexitProgram *v1alpha1.FexitProgram, opts v1.UpdateOptions) (result *v1alpha1.FexitProgram, err error) {
	result = &v1alpha1.FexitProgram{}
	err = c.client.Put().
		Resource("fexitprograms").
		Name(fexitProgram.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fexitProgram).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the fexitProgram and deletes it. Returns an error if one occurs.
func (c *fexitPrograms) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("fexitprograms").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *fexitPrograms) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("fexitprograms").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched fexitProgram.
func (c *fexitPrograms) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FexitProgram, err error) {
	result = &v1alpha1.FexitProgram{}
	err = c.client.Patch(pt).
		Resource("fexitprograms").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}