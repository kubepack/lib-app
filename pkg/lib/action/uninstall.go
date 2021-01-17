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

package action

import (
	"time"

	"kubepack.dev/lib-app/pkg/action"

	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/release"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type UninstallOptions struct {
	DisableHooks bool          `json:"disableHooks"`
	DryRun       bool          `json:"dryRun"`
	KeepHistory  bool          `json:"keepHistory"`
	Timeout      time.Duration `json:"timeout"`
}

type Uninstaller struct {
	cfg *action.Configuration

	opts        UninstallOptions
	releaseName string
	result      *release.UninstallReleaseResponse
}

func NewUninstaller(getter genericclioptions.RESTClientGetter, namespace string, helmDriver string) (*Uninstaller, error) {
	cfg := new(action.Configuration)
	// TODO: Use secret driver for which namespace?
	err := cfg.Init(getter, namespace, helmDriver, debug)
	if err != nil {
		return nil, err
	}
	cfg.Capabilities = chartutil.DefaultCapabilities

	return NewUninstallerForConfig(cfg), nil
}

func NewUninstallerForConfig(cfg *action.Configuration) *Uninstaller {
	return &Uninstaller{
		cfg: cfg,
	}
}

func (x *Uninstaller) WithOptions(opts UninstallOptions) *Uninstaller {
	x.opts = opts
	return x
}

func (x *Uninstaller) WithReleaseName(name string) *Uninstaller {
	x.releaseName = name
	return x
}

func (x *Uninstaller) Run() (*release.UninstallReleaseResponse, error) {
	cmd := action.NewUninstall(x.cfg)
	cmd.DisableHooks = x.opts.DisableHooks
	cmd.DryRun = x.opts.DryRun
	cmd.KeepHistory = x.opts.KeepHistory
	cmd.Timeout = x.opts.Timeout

	return cmd.Run(x.releaseName)
}

func (x *Uninstaller) Do() error {
	var err error
	x.result, err = x.Run()
	return err
}

func (x *Uninstaller) Result() *release.UninstallReleaseResponse {
	return x.result
}
