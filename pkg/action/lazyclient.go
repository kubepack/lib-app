/*
Copyright The Helm Authors.

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

package action

import (
	"context"
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	appapi "sigs.k8s.io/application/api/app/v1beta1"
	appcs "sigs.k8s.io/application/client/clientset/versioned"
	"sigs.k8s.io/application/client/clientset/versioned/typed/app/v1beta1"
)

// lazyClient is a workaround to deal with Kubernetes having an unstable client API.
// In Kubernetes v1.18 the defaults where removed which broke creating a
// client without an explicit configuration. ಠ_ಠ
type lazyClient struct {
	// client caches an initialized kubernetes client
	initClient sync.Once
	client     kubernetes.Interface
	appClient  appcs.Interface
	clientErr  error

	// clientFn loads a kubernetes client
	clientFn    func() (*kubernetes.Clientset, error)
	appClientFn func() (*appcs.Clientset, error)

	// namespace passed to each client request
	namespace string
}

func (s *lazyClient) init() error {
	s.initClient.Do(func() {
		s.client, s.clientErr = s.clientFn()
		s.appClient, s.clientErr = s.appClientFn()
	})
	return s.clientErr
}

// applicationClient implements a coreappv1beta1.ApplicationInterface
type applicationClient struct{ *lazyClient }

var _ v1beta1.ApplicationInterface = (*applicationClient)(nil)

func newApplicationClient(lc *lazyClient) *applicationClient {
	return &applicationClient{lazyClient: lc}
}

func (c *applicationClient) Create(ctx context.Context, application *appapi.Application, opts metav1.CreateOptions) (*appapi.Application, error) {
	if err := c.init(); err != nil {
		return nil, err
	}
	return c.appClient.AppV1beta1().Applications(c.namespace).Create(ctx, application, opts)
}

func (c *applicationClient) Update(ctx context.Context, application *appapi.Application, opts metav1.UpdateOptions) (*appapi.Application, error) {
	if err := c.init(); err != nil {
		return nil, err
	}
	return c.appClient.AppV1beta1().Applications(c.namespace).Update(ctx, application, opts)
}

func (c *applicationClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	if err := c.init(); err != nil {
		return err
	}
	return c.appClient.AppV1beta1().Applications(c.namespace).Delete(ctx, name, opts)
}

func (c *applicationClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	if err := c.init(); err != nil {
		return err
	}
	return c.appClient.AppV1beta1().Applications(c.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (c *applicationClient) Get(ctx context.Context, name string, opts metav1.GetOptions) (*appapi.Application, error) {
	if err := c.init(); err != nil {
		return nil, err
	}
	return c.appClient.AppV1beta1().Applications(c.namespace).Get(ctx, name, opts)
}

func (c *applicationClient) List(ctx context.Context, opts metav1.ListOptions) (*appapi.ApplicationList, error) {
	if err := c.init(); err != nil {
		return nil, err
	}
	return c.appClient.AppV1beta1().Applications(c.namespace).List(ctx, opts)
}

func (c *applicationClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	if err := c.init(); err != nil {
		return nil, err
	}
	return c.appClient.AppV1beta1().Applications(c.namespace).Watch(ctx, opts)
}

func (c *applicationClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*appapi.Application, error) {
	if err := c.init(); err != nil {
		return nil, err
	}
	return c.appClient.AppV1beta1().Applications(c.namespace).Patch(ctx, name, pt, data, opts, subresources...)
}

func (c *applicationClient) UpdateStatus(ctx context.Context, application *appapi.Application, opts metav1.UpdateOptions) (*appapi.Application, error) {
	if err := c.init(); err != nil {
		return nil, err
	}
	return c.appClient.AppV1beta1().Applications(c.namespace).UpdateStatus(ctx, application, opts)
}
