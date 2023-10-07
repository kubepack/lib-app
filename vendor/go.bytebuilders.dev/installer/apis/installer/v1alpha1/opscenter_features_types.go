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

	Release         ReleaseInfo                `json:"release"`
	Repositories    map[string]*HelmRepository `json:"repositories"`
	Registry        RegistryInfo               `json:"registry"`
	ClusterManagers []string                   `json:"clusterManagers"`
	CAPI            CapiPresetsSpec            `json:"capi"`
}

type ReleaseInfo struct {
	// +kubebuilder:default:=dev
	Channel ReleaseChannel `json:"channel"`
}

// +kubebuilder:validation:Enum=stable;testing;dev
type ReleaseChannel string

const (
	ChannelStable  ReleaseChannel = "stable"
	ChannelTesting ReleaseChannel = "testing"
	ChannelDev     ReleaseChannel = "dev"
)

type HelmRepository struct {
	// URL of the Helm repository, a valid URL contains at least a protocol and
	// host.
	// +required
	URL string `json:"url"`

	// SecretRef specifies the Secret containing authentication credentials
	// for the HelmRepository.
	// For HTTP/S basic auth the secret must contain 'username' and 'password'
	// fields.
	// For TLS the secret must contain a 'certFile' and 'keyFile', and/or
	// 'caFile' fields.
	// +optional
	SecretName string `json:"secretName,omitempty"`

	// Interval at which to check the URL for updates.
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern="^([0-9]+(\\.[0-9]+)?(ms|s|m|h))+$"
	// +required
	Interval metav1.Duration `json:"interval"`

	// Type of the HelmRepository.
	// When this field is set to  "oci", the URL field value must be prefixed with "oci://".
	// +kubebuilder:validation:Enum=default;oci
	// +optional
	Type string `json:"type,omitempty"`

	// Provider used for authentication, can be 'aws', 'azure', 'gcp' or 'generic'.
	// This field is optional, and only taken into account if the .spec.type field is set to 'oci'.
	// When not specified, defaults to 'generic'.
	// +kubebuilder:validation:Enum=generic;aws;azure;gcp
	// +kubebuilder:default:=generic
	// +optional
	Provider string `json:"provider,omitempty"`
}

type RepositoryCredential map[string]string

type RegistryInfo struct {
	Credentials RepositoryCredential `json:"credentials"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpscenterFeaturesList is a list of OpscenterFeaturess
type OpscenterFeaturesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpscenterFeatures CRD objects
	Items []OpscenterFeatures `json:"items,omitempty"`
}
