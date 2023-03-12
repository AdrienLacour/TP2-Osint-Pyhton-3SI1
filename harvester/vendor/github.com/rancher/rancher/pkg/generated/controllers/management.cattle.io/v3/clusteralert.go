/*
Copyright 2022 Rancher Labs, Inc.

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

	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
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

type ClusterAlertHandler func(string, *v3.ClusterAlert) (*v3.ClusterAlert, error)

type ClusterAlertController interface {
	generic.ControllerMeta
	ClusterAlertClient

	OnChange(ctx context.Context, name string, sync ClusterAlertHandler)
	OnRemove(ctx context.Context, name string, sync ClusterAlertHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() ClusterAlertCache
}

type ClusterAlertClient interface {
	Create(*v3.ClusterAlert) (*v3.ClusterAlert, error)
	Update(*v3.ClusterAlert) (*v3.ClusterAlert, error)
	UpdateStatus(*v3.ClusterAlert) (*v3.ClusterAlert, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v3.ClusterAlert, error)
	List(namespace string, opts metav1.ListOptions) (*v3.ClusterAlertList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.ClusterAlert, err error)
}

type ClusterAlertCache interface {
	Get(namespace, name string) (*v3.ClusterAlert, error)
	List(namespace string, selector labels.Selector) ([]*v3.ClusterAlert, error)

	AddIndexer(indexName string, indexer ClusterAlertIndexer)
	GetByIndex(indexName, key string) ([]*v3.ClusterAlert, error)
}

type ClusterAlertIndexer func(obj *v3.ClusterAlert) ([]string, error)

type clusterAlertController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewClusterAlertController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) ClusterAlertController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &clusterAlertController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromClusterAlertHandlerToHandler(sync ClusterAlertHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.ClusterAlert
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.ClusterAlert))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *clusterAlertController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.ClusterAlert))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateClusterAlertDeepCopyOnChange(client ClusterAlertClient, obj *v3.ClusterAlert, handler func(obj *v3.ClusterAlert) (*v3.ClusterAlert, error)) (*v3.ClusterAlert, error) {
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

func (c *clusterAlertController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *clusterAlertController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *clusterAlertController) OnChange(ctx context.Context, name string, sync ClusterAlertHandler) {
	c.AddGenericHandler(ctx, name, FromClusterAlertHandlerToHandler(sync))
}

func (c *clusterAlertController) OnRemove(ctx context.Context, name string, sync ClusterAlertHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromClusterAlertHandlerToHandler(sync)))
}

func (c *clusterAlertController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *clusterAlertController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *clusterAlertController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *clusterAlertController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *clusterAlertController) Cache() ClusterAlertCache {
	return &clusterAlertCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *clusterAlertController) Create(obj *v3.ClusterAlert) (*v3.ClusterAlert, error) {
	result := &v3.ClusterAlert{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *clusterAlertController) Update(obj *v3.ClusterAlert) (*v3.ClusterAlert, error) {
	result := &v3.ClusterAlert{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *clusterAlertController) UpdateStatus(obj *v3.ClusterAlert) (*v3.ClusterAlert, error) {
	result := &v3.ClusterAlert{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *clusterAlertController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *clusterAlertController) Get(namespace, name string, options metav1.GetOptions) (*v3.ClusterAlert, error) {
	result := &v3.ClusterAlert{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *clusterAlertController) List(namespace string, opts metav1.ListOptions) (*v3.ClusterAlertList, error) {
	result := &v3.ClusterAlertList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *clusterAlertController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *clusterAlertController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v3.ClusterAlert, error) {
	result := &v3.ClusterAlert{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type clusterAlertCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *clusterAlertCache) Get(namespace, name string) (*v3.ClusterAlert, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.ClusterAlert), nil
}

func (c *clusterAlertCache) List(namespace string, selector labels.Selector) (ret []*v3.ClusterAlert, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.ClusterAlert))
	})

	return ret, err
}

func (c *clusterAlertCache) AddIndexer(indexName string, indexer ClusterAlertIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.ClusterAlert))
		},
	}))
}

func (c *clusterAlertCache) GetByIndex(indexName, key string) (result []*v3.ClusterAlert, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.ClusterAlert, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.ClusterAlert))
	}
	return result, nil
}

type ClusterAlertStatusHandler func(obj *v3.ClusterAlert, status v3.AlertStatus) (v3.AlertStatus, error)

type ClusterAlertGeneratingHandler func(obj *v3.ClusterAlert, status v3.AlertStatus) ([]runtime.Object, v3.AlertStatus, error)

func RegisterClusterAlertStatusHandler(ctx context.Context, controller ClusterAlertController, condition condition.Cond, name string, handler ClusterAlertStatusHandler) {
	statusHandler := &clusterAlertStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromClusterAlertHandlerToHandler(statusHandler.sync))
}

func RegisterClusterAlertGeneratingHandler(ctx context.Context, controller ClusterAlertController, apply apply.Apply,
	condition condition.Cond, name string, handler ClusterAlertGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &clusterAlertGeneratingHandler{
		ClusterAlertGeneratingHandler: handler,
		apply:                         apply,
		name:                          name,
		gvk:                           controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterClusterAlertStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type clusterAlertStatusHandler struct {
	client    ClusterAlertClient
	condition condition.Cond
	handler   ClusterAlertStatusHandler
}

func (a *clusterAlertStatusHandler) sync(key string, obj *v3.ClusterAlert) (*v3.ClusterAlert, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type clusterAlertGeneratingHandler struct {
	ClusterAlertGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *clusterAlertGeneratingHandler) Remove(key string, obj *v3.ClusterAlert) (*v3.ClusterAlert, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v3.ClusterAlert{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *clusterAlertGeneratingHandler) Handle(obj *v3.ClusterAlert, status v3.AlertStatus) (v3.AlertStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.ClusterAlertGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
