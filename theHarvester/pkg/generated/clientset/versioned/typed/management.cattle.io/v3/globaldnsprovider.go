/*
Copyright 2023 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GlobalDnsProvidersGetter has a method to return a GlobalDnsProviderInterface.
// A group's client should implement this interface.
type GlobalDnsProvidersGetter interface {
	GlobalDnsProviders(namespace string) GlobalDnsProviderInterface
}

// GlobalDnsProviderInterface has methods to work with GlobalDnsProvider resources.
type GlobalDnsProviderInterface interface {
	Create(ctx context.Context, globalDnsProvider *v3.GlobalDnsProvider, opts v1.CreateOptions) (*v3.GlobalDnsProvider, error)
	Update(ctx context.Context, globalDnsProvider *v3.GlobalDnsProvider, opts v1.UpdateOptions) (*v3.GlobalDnsProvider, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.GlobalDnsProvider, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.GlobalDnsProviderList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.GlobalDnsProvider, err error)
	GlobalDnsProviderExpansion
}

// globalDnsProviders implements GlobalDnsProviderInterface
type globalDnsProviders struct {
	client rest.Interface
	ns     string
}

// newGlobalDnsProviders returns a GlobalDnsProviders
func newGlobalDnsProviders(c *ManagementV3Client, namespace string) *globalDnsProviders {
	return &globalDnsProviders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the globalDnsProvider, and returns the corresponding globalDnsProvider object, and an error if there is any.
func (c *globalDnsProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.GlobalDnsProvider, err error) {
	result = &v3.GlobalDnsProvider{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("globaldnsproviders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GlobalDnsProviders that match those selectors.
func (c *globalDnsProviders) List(ctx context.Context, opts v1.ListOptions) (result *v3.GlobalDnsProviderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.GlobalDnsProviderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("globaldnsproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested globalDnsProviders.
func (c *globalDnsProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("globaldnsproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a globalDnsProvider and creates it.  Returns the server's representation of the globalDnsProvider, and an error, if there is any.
func (c *globalDnsProviders) Create(ctx context.Context, globalDnsProvider *v3.GlobalDnsProvider, opts v1.CreateOptions) (result *v3.GlobalDnsProvider, err error) {
	result = &v3.GlobalDnsProvider{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("globaldnsproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalDnsProvider).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a globalDnsProvider and updates it. Returns the server's representation of the globalDnsProvider, and an error, if there is any.
func (c *globalDnsProviders) Update(ctx context.Context, globalDnsProvider *v3.GlobalDnsProvider, opts v1.UpdateOptions) (result *v3.GlobalDnsProvider, err error) {
	result = &v3.GlobalDnsProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("globaldnsproviders").
		Name(globalDnsProvider.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalDnsProvider).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the globalDnsProvider and deletes it. Returns an error if one occurs.
func (c *globalDnsProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("globaldnsproviders").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *globalDnsProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("globaldnsproviders").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched globalDnsProvider.
func (c *globalDnsProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.GlobalDnsProvider, err error) {
	result = &v3.GlobalDnsProvider{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("globaldnsproviders").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
