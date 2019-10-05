// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	operatorv1 "github.com/openshift/api/operator/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConsoles implements ConsoleInterface
type FakeConsoles struct {
	Fake *FakeOperatorV1
}

var consolesResource = schema.GroupVersionResource{Group: "operator.openshift.io", Version: "v1", Resource: "consoles"}

var consolesKind = schema.GroupVersionKind{Group: "operator.openshift.io", Version: "v1", Kind: "Console"}

// Get takes name of the console, and returns the corresponding console object, and an error if there is any.
func (c *FakeConsoles) Get(name string, options v1.GetOptions) (result *operatorv1.Console, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(consolesResource, name), &operatorv1.Console{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Console), err
}

// List takes label and field selectors, and returns the list of Consoles that match those selectors.
func (c *FakeConsoles) List(opts v1.ListOptions) (result *operatorv1.ConsoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(consolesResource, consolesKind, opts), &operatorv1.ConsoleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorv1.ConsoleList{ListMeta: obj.(*operatorv1.ConsoleList).ListMeta}
	for _, item := range obj.(*operatorv1.ConsoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested consoles.
func (c *FakeConsoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(consolesResource, opts))
}

// Create takes the representation of a console and creates it.  Returns the server's representation of the console, and an error, if there is any.
func (c *FakeConsoles) Create(console *operatorv1.Console) (result *operatorv1.Console, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(consolesResource, console), &operatorv1.Console{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Console), err
}

// Update takes the representation of a console and updates it. Returns the server's representation of the console, and an error, if there is any.
func (c *FakeConsoles) Update(console *operatorv1.Console) (result *operatorv1.Console, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(consolesResource, console), &operatorv1.Console{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Console), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConsoles) UpdateStatus(console *operatorv1.Console) (*operatorv1.Console, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(consolesResource, "status", console), &operatorv1.Console{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Console), err
}

// Delete takes name of the console and deletes it. Returns an error if one occurs.
func (c *FakeConsoles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(consolesResource, name), &operatorv1.Console{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConsoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(consolesResource, listOptions)

	_, err := c.Fake.Invokes(action, &operatorv1.ConsoleList{})
	return err
}

// Patch applies the patch and returns the patched console.
func (c *FakeConsoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *operatorv1.Console, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(consolesResource, name, pt, data, subresources...), &operatorv1.Console{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Console), err
}