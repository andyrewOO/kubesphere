/*
Copyright 2019 The KubeSphere authors.

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
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubesphere.io/kubesphere/pkg/apis/devops/v1alpha1"
	scheme "kubesphere.io/kubesphere/pkg/client/clientset/versioned/scheme"
)

// S2iRunsGetter has a method to return a S2iRunInterface.
// A group's client should implement this interface.
type S2iRunsGetter interface {
	S2iRuns(namespace string) S2iRunInterface
}

// S2iRunInterface has methods to work with S2iRun resources.
type S2iRunInterface interface {
	Create(*v1alpha1.S2iRun) (*v1alpha1.S2iRun, error)
	Update(*v1alpha1.S2iRun) (*v1alpha1.S2iRun, error)
	UpdateStatus(*v1alpha1.S2iRun) (*v1alpha1.S2iRun, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.S2iRun, error)
	List(opts v1.ListOptions) (*v1alpha1.S2iRunList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.S2iRun, err error)
	S2iRunExpansion
}

// s2iRuns implements S2iRunInterface
type s2iRuns struct {
	client rest.Interface
	ns     string
}

// newS2iRuns returns a S2iRuns
func newS2iRuns(c *DevopsV1alpha1Client, namespace string) *s2iRuns {
	return &s2iRuns{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the s2iRun, and returns the corresponding s2iRun object, and an error if there is any.
func (c *s2iRuns) Get(name string, options v1.GetOptions) (result *v1alpha1.S2iRun, err error) {
	result = &v1alpha1.S2iRun{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s2iruns").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of S2iRuns that match those selectors.
func (c *s2iRuns) List(opts v1.ListOptions) (result *v1alpha1.S2iRunList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.S2iRunList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s2iruns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested s2iRuns.
func (c *s2iRuns) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("s2iruns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a s2iRun and creates it.  Returns the server's representation of the s2iRun, and an error, if there is any.
func (c *s2iRuns) Create(s2iRun *v1alpha1.S2iRun) (result *v1alpha1.S2iRun, err error) {
	result = &v1alpha1.S2iRun{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("s2iruns").
		Body(s2iRun).
		Do().
		Into(result)
	return
}

// Update takes the representation of a s2iRun and updates it. Returns the server's representation of the s2iRun, and an error, if there is any.
func (c *s2iRuns) Update(s2iRun *v1alpha1.S2iRun) (result *v1alpha1.S2iRun, err error) {
	result = &v1alpha1.S2iRun{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s2iruns").
		Name(s2iRun.Name).
		Body(s2iRun).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *s2iRuns) UpdateStatus(s2iRun *v1alpha1.S2iRun) (result *v1alpha1.S2iRun, err error) {
	result = &v1alpha1.S2iRun{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s2iruns").
		Name(s2iRun.Name).
		SubResource("status").
		Body(s2iRun).
		Do().
		Into(result)
	return
}

// Delete takes name of the s2iRun and deletes it. Returns an error if one occurs.
func (c *s2iRuns) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s2iruns").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *s2iRuns) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s2iruns").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched s2iRun.
func (c *s2iRuns) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.S2iRun, err error) {
	result = &v1alpha1.S2iRun{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("s2iruns").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
