/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kmodules.xyz/resource-metadata/apis/shared"
)

const (
	ResourceKindOpscenterFeatures = "OpscenterFeatures"
	ResourceOpscenterFeatures     = "opscenterfeatures"
	ResourceOpscenterFeaturess    = "opscenterfeaturess"
)

// OpscenterFeatures defines the schama for OpscenterFeatures Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=opscenterfeaturess,singular=opscenterfeatures,categories={kubeops,appscode}
type OpscenterFeatures struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpscenterFeaturesSpec `json:"spec,omitempty"`
}

// OpscenterFeaturesSpec is the schema for OpscenterFeatures Operator values file
type OpscenterFeaturesSpec struct {
	NameOverride     string `json:"nameOverride"`
	FullnameOverride string `json:"fullnameOverride"`

	// +optional
	OfflineInstaller bool                     `json:"offlineInstaller"`
	Image            shared.ImageRegistrySpec `json:"image"`
	Helm             OpscenterHelmSpec        `json:"helm"`
	Registry         shared.RegistryInfo      `json:"registry"`
	ClusterMetadata  ClusterMetadata          `json:"clusterMetadata"`
}

type ClusterMetadata struct {
	Uid             string   `json:"uid"`
	Name            string   `json:"name"`
	ClusterManagers []string `json:"clusterManagers"`
	// +optional
	CAPI CapiMetadata `json:"capi"`
}

type CapiMetadata struct {
	// +optional
	Provider  CAPIProvider `json:"provider"`
	Namespace string       `json:"namespace"`
}

// +kubebuilder:validation:Enum=capa;capg;capz
type CAPIProvider string

const (
	CAPIProviderDisabled CAPIProvider = ""
	CAPIProviderCAPA     CAPIProvider = "capa"
	CAPIProviderCAPG     CAPIProvider = "capg"
	CAPIProviderCAPZ     CAPIProvider = "capz"
)

type OpscenterHelmSpec struct {
	CreateNamespace bool                              `json:"createNamespace"`
	Repositories    map[string]*shared.HelmRepository `json:"repositories"`
	Releases        map[string]*HelmRelease           `json:"releases"`
}

type HelmRelease struct {
	Version string `json:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpscenterFeaturesList is a list of OpscenterFeaturess
type OpscenterFeaturesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpscenterFeatures CRD objects
	Items []OpscenterFeatures `json:"items,omitempty"`
}
