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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type SettingHandler func(string, *v1beta1.Setting) (*v1beta1.Setting, error)

type SettingController interface {
	generic.ControllerMeta
	SettingClient

	OnChange(ctx context.Context, name string, sync SettingHandler)
	OnRemove(ctx context.Context, name string, sync SettingHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() SettingCache
}

type SettingClient interface {
	Create(*v1beta1.Setting) (*v1beta1.Setting, error)
	Update(*v1beta1.Setting) (*v1beta1.Setting, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1beta1.Setting, error)
	List(namespace string, opts metav1.ListOptions) (*v1beta1.SettingList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Setting, err error)
}

type SettingCache interface {
	Get(namespace, name string) (*v1beta1.Setting, error)
	List(namespace string, selector labels.Selector) ([]*v1beta1.Setting, error)

	AddIndexer(indexName string, indexer SettingIndexer)
	GetByIndex(indexName, key string) ([]*v1beta1.Setting, error)
}

type SettingIndexer func(obj *v1beta1.Setting) ([]string, error)

type settingController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewSettingController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) SettingController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &settingController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromSettingHandlerToHandler(sync SettingHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1beta1.Setting
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1beta1.Setting))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *settingController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1beta1.Setting))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateSettingDeepCopyOnChange(client SettingClient, obj *v1beta1.Setting, handler func(obj *v1beta1.Setting) (*v1beta1.Setting, error)) (*v1beta1.Setting, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *settingController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *settingController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *settingController) OnChange(ctx context.Context, name string, sync SettingHandler) {
	c.AddGenericHandler(ctx, name, FromSettingHandlerToHandler(sync))
}

func (c *settingController) OnRemove(ctx context.Context, name string, sync SettingHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromSettingHandlerToHandler(sync)))
}

func (c *settingController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *settingController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *settingController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *settingController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *settingController) Cache() SettingCache {
	return &settingCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *settingController) Create(obj *v1beta1.Setting) (*v1beta1.Setting, error) {
	result := &v1beta1.Setting{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *settingController) Update(obj *v1beta1.Setting) (*v1beta1.Setting, error) {
	result := &v1beta1.Setting{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *settingController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *settingController) Get(namespace, name string, options metav1.GetOptions) (*v1beta1.Setting, error) {
	result := &v1beta1.Setting{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *settingController) List(namespace string, opts metav1.ListOptions) (*v1beta1.SettingList, error) {
	result := &v1beta1.SettingList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *settingController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *settingController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1beta1.Setting, error) {
	result := &v1beta1.Setting{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type settingCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *settingCache) Get(namespace, name string) (*v1beta1.Setting, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1beta1.Setting), nil
}

func (c *settingCache) List(namespace string, selector labels.Selector) (ret []*v1beta1.Setting, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Setting))
	})

	return ret, err
}

func (c *settingCache) AddIndexer(indexName string, indexer SettingIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1beta1.Setting))
		},
	}))
}

func (c *settingCache) GetByIndex(indexName, key string) (result []*v1beta1.Setting, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1beta1.Setting, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1beta1.Setting))
	}
	return result, nil
}
