/*
Copyright AppsCode Inc. and Contributors

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

package fusion

import (
	"context"
	"reflect"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type FakeClient struct {
	store map[reflect.Type]ObjStore
}

type ObjStore map[client.ObjectKey]client.Object

var _ client.Reader = &FakeClient{}

func NewFakeClient(list ...client.Object) client.Reader {
	f := FakeClient{
		store: make(map[reflect.Type]ObjStore),
	}
	for _, obj := range list {
		typ := reflect.TypeOf(obj)

		store, found := f.store[typ]
		if !found {
			store = make(ObjStore)
		}
		store[client.ObjectKeyFromObject(obj)] = obj
		f.store[typ] = store
	}
	return &f
}

func (f *FakeClient) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	typ := reflect.TypeOf(obj)
	store, found := f.store[typ]
	if !found {
		return apierrors.NewNotFound(schema.GroupResource{
			Group:    "",
			Resource: "",
		}, key.Name)
	}
	result, exists := store[key]
	if !exists {
		return apierrors.NewNotFound(schema.GroupResource{
			Group:    "",
			Resource: "",
		}, key.Name)
	}
	assign(obj, result)
	return nil
}

func assign(target, src any) {
	srcValue := reflect.ValueOf(src)
	if srcValue.Kind() == reflect.Pointer {
		srcValue = srcValue.Elem()
	}
	reflect.ValueOf(target).Elem().Set(srcValue)
}

func (f *FakeClient) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	// TODO implement me
	panic("implement me")
}
