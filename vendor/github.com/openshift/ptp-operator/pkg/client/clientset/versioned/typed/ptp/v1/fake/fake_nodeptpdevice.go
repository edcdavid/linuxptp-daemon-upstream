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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	ptpv1 "github.com/openshift/ptp-operator/pkg/apis/ptp/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodePtpDevices implements NodePtpDeviceInterface
type FakeNodePtpDevices struct {
	Fake *FakePtpV1
	ns   string
}

var nodeptpdevicesResource = schema.GroupVersionResource{Group: "ptp.openshift.io", Version: "v1", Resource: "nodeptpdevices"}

var nodeptpdevicesKind = schema.GroupVersionKind{Group: "ptp.openshift.io", Version: "v1", Kind: "NodePtpDevice"}

// Get takes name of the nodePtpDevice, and returns the corresponding nodePtpDevice object, and an error if there is any.
func (c *FakeNodePtpDevices) Get(name string, options v1.GetOptions) (result *ptpv1.NodePtpDevice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nodeptpdevicesResource, c.ns, name), &ptpv1.NodePtpDevice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ptpv1.NodePtpDevice), err
}

// List takes label and field selectors, and returns the list of NodePtpDevices that match those selectors.
func (c *FakeNodePtpDevices) List(opts v1.ListOptions) (result *ptpv1.NodePtpDeviceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nodeptpdevicesResource, nodeptpdevicesKind, c.ns, opts), &ptpv1.NodePtpDeviceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &ptpv1.NodePtpDeviceList{ListMeta: obj.(*ptpv1.NodePtpDeviceList).ListMeta}
	for _, item := range obj.(*ptpv1.NodePtpDeviceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodePtpDevices.
func (c *FakeNodePtpDevices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nodeptpdevicesResource, c.ns, opts))

}

// Create takes the representation of a nodePtpDevice and creates it.  Returns the server's representation of the nodePtpDevice, and an error, if there is any.
func (c *FakeNodePtpDevices) Create(nodePtpDevice *ptpv1.NodePtpDevice) (result *ptpv1.NodePtpDevice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nodeptpdevicesResource, c.ns, nodePtpDevice), &ptpv1.NodePtpDevice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ptpv1.NodePtpDevice), err
}

// Update takes the representation of a nodePtpDevice and updates it. Returns the server's representation of the nodePtpDevice, and an error, if there is any.
func (c *FakeNodePtpDevices) Update(nodePtpDevice *ptpv1.NodePtpDevice) (result *ptpv1.NodePtpDevice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nodeptpdevicesResource, c.ns, nodePtpDevice), &ptpv1.NodePtpDevice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ptpv1.NodePtpDevice), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodePtpDevices) UpdateStatus(nodePtpDevice *ptpv1.NodePtpDevice) (*ptpv1.NodePtpDevice, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(nodeptpdevicesResource, "status", c.ns, nodePtpDevice), &ptpv1.NodePtpDevice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ptpv1.NodePtpDevice), err
}

// Delete takes name of the nodePtpDevice and deletes it. Returns an error if one occurs.
func (c *FakeNodePtpDevices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nodeptpdevicesResource, c.ns, name), &ptpv1.NodePtpDevice{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodePtpDevices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nodeptpdevicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &ptpv1.NodePtpDeviceList{})
	return err
}

// Patch applies the patch and returns the patched nodePtpDevice.
func (c *FakeNodePtpDevices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *ptpv1.NodePtpDevice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nodeptpdevicesResource, c.ns, name, pt, data, subresources...), &ptpv1.NodePtpDevice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ptpv1.NodePtpDevice), err
}
