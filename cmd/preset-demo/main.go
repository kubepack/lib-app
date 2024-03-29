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

package main

import (
	"fmt"
	"path/filepath"

	"kubepack.dev/lib-helm/pkg/action"
	actionx "kubepack.dev/lib-helm/pkg/action"
	"kubepack.dev/lib-helm/pkg/repo"
	"kubepack.dev/lib-helm/pkg/values"

	flag "github.com/spf13/pflag"
	"gomodules.xyz/x/crypto/rand"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
	clientcmdutil "kmodules.xyz/client-go/tools/clientcmd"
	chartsapi "x-helm.dev/apimachinery/apis/charts/v1alpha1"
	releasesapi "x-helm.dev/apimachinery/apis/releases/v1alpha1"
)

var HelmRegistry = repo.NewDiskCacheRegistry()

func main() {
	var (
		masterURL      = ""
		kubeconfigPath = filepath.Join(homedir.HomeDir(), ".kube", "config")

		url     = "https://raw.githubusercontent.com/kubepack/preset-testdata/master/stable/"
		name    = "hello"
		version = "0.1.0"
	)
	flag.StringVar(&masterURL, "master", masterURL, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	flag.StringVar(&kubeconfigPath, "kubeconfig", kubeconfigPath, "Path to kubeconfig file with authorization information (the master location is set by the master flag).")
	flag.StringVar(&url, "url", url, "Chart repo url")
	flag.StringVar(&name, "name", name, "Name of bundle")
	flag.StringVar(&version, "version", version, "Version of bundle")
	flag.Parse()

	cc := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: masterURL}})
	kubeconfig, err := cc.RawConfig()
	if err != nil {
		klog.Fatal(err)
	}
	getter := clientcmdutil.NewClientGetter(&kubeconfig)

	ref := chartsapi.ChartPresetFlatRef{
		ChartSourceFlatRef: releasesapi.ChartSourceFlatRef{
			Name:            name,
			Version:         version,
			SourceAPIGroup:  releasesapi.SourceGroupLegacy,
			SourceKind:      releasesapi.SourceKindLegacy,
			SourceNamespace: "",
			SourceName:      url,
		},
		// PresetGroup:    chartsapi.GroupVersion.Group,
		// PresetKind:     chartsapi.ResourceKindClusterChartPreset,
		Variant: "unified",
		// PresetSelector: "",
		Namespace: "default",
	}
	// encoder := form.NewEncoder()
	// encoder.SetTagName("json")
	// qv, err := encoder.Encode(&ref)
	// fmt.Println(qv.Encode())

	if err := DD(getter, ref); err != nil {
		klog.Fatalln(err)
	}
}

func DD(getter genericclioptions.RESTClientGetter, ref chartsapi.ChartPresetFlatRef) error {
	kc, err := actionx.NewUncachedClient(getter)
	if err != nil {
		return err
	}

	chrt, err := HelmRegistry.GetChart(ref.ChartSourceFlatRef.ToAPIObject())
	if err != nil {
		return err
	}

	vals, err := values.MergePresetValues(kc, chrt.Chart, ref)
	if err != nil {
		return err
	}

	i, err := action.NewInstaller(getter, ref.Namespace, "secret")
	if err != nil {
		return err
	}
	i.WithRegistry(HelmRegistry).
		WithOptions(action.InstallOptions{
			ChartSourceFlatRef: ref.ChartSourceFlatRef,
			Options: values.Options{
				ReplaceValues: vals,
			},
			DryRun:       false,
			DisableHooks: false,
			Replace:      false,
			Wait:         false,
			Devel:        false,
			Timeout:      0,
			Namespace:    ref.Namespace,
			ReleaseName:  rand.WithUniqSuffix(ref.Name),
			Atomic:       false,
			SkipCRDs:     false,
		})
	rel, err := i.Run()
	if err != nil {
		return err
	}
	fmt.Println(rel)
	return nil
}
