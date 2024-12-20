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

package fake

import (
	"context"

	v1alpha1 "github.com/bpfman/bpfman-operator/apis/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeXdpNsPrograms implements XdpNsProgramInterface
type FakeXdpNsPrograms struct {
	Fake *FakeBpfmanV1alpha1
	ns   string
}

var xdpnsprogramsResource = v1alpha1.SchemeGroupVersion.WithResource("xdpnsprograms")

var xdpnsprogramsKind = v1alpha1.SchemeGroupVersion.WithKind("XdpNsProgram")

// Get takes name of the xdpNsProgram, and returns the corresponding xdpNsProgram object, and an error if there is any.
func (c *FakeXdpNsPrograms) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.XdpNsProgram, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(xdpnsprogramsResource, c.ns, name), &v1alpha1.XdpNsProgram{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.XdpNsProgram), err
}

// List takes label and field selectors, and returns the list of XdpNsPrograms that match those selectors.
func (c *FakeXdpNsPrograms) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.XdpNsProgramList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(xdpnsprogramsResource, xdpnsprogramsKind, c.ns, opts), &v1alpha1.XdpNsProgramList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.XdpNsProgramList{ListMeta: obj.(*v1alpha1.XdpNsProgramList).ListMeta}
	for _, item := range obj.(*v1alpha1.XdpNsProgramList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested xdpNsPrograms.
func (c *FakeXdpNsPrograms) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(xdpnsprogramsResource, c.ns, opts))

}

// Create takes the representation of a xdpNsProgram and creates it.  Returns the server's representation of the xdpNsProgram, and an error, if there is any.
func (c *FakeXdpNsPrograms) Create(ctx context.Context, xdpNsProgram *v1alpha1.XdpNsProgram, opts v1.CreateOptions) (result *v1alpha1.XdpNsProgram, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(xdpnsprogramsResource, c.ns, xdpNsProgram), &v1alpha1.XdpNsProgram{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.XdpNsProgram), err
}

// Update takes the representation of a xdpNsProgram and updates it. Returns the server's representation of the xdpNsProgram, and an error, if there is any.
func (c *FakeXdpNsPrograms) Update(ctx context.Context, xdpNsProgram *v1alpha1.XdpNsProgram, opts v1.UpdateOptions) (result *v1alpha1.XdpNsProgram, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(xdpnsprogramsResource, c.ns, xdpNsProgram), &v1alpha1.XdpNsProgram{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.XdpNsProgram), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeXdpNsPrograms) UpdateStatus(ctx context.Context, xdpNsProgram *v1alpha1.XdpNsProgram, opts v1.UpdateOptions) (*v1alpha1.XdpNsProgram, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(xdpnsprogramsResource, "status", c.ns, xdpNsProgram), &v1alpha1.XdpNsProgram{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.XdpNsProgram), err
}

// Delete takes name of the xdpNsProgram and deletes it. Returns an error if one occurs.
func (c *FakeXdpNsPrograms) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(xdpnsprogramsResource, c.ns, name, opts), &v1alpha1.XdpNsProgram{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeXdpNsPrograms) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(xdpnsprogramsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.XdpNsProgramList{})
	return err
}

// Patch applies the patch and returns the patched xdpNsProgram.
func (c *FakeXdpNsPrograms) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.XdpNsProgram, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(xdpnsprogramsResource, c.ns, name, pt, data, subresources...), &v1alpha1.XdpNsProgram{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.XdpNsProgram), err
}
