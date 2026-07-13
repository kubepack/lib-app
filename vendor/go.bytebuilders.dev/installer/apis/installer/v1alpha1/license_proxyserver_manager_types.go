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
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kmodules.xyz/resource-metadata/apis/shared"
)

const (
	ResourceKindLicenseProxyserverManager = "LicenseProxyserverManager"
	ResourceLicenseProxyserverManager     = "licenseproxyservermanager"
	ResourceLicenseProxyserverManagers    = "licenseproxyservermanagers"
)

// LicenseProxyserverManager defines the schama for LicenseProxyserverManager Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=licenseproxyservermanagers,singular=licenseproxyservermanager,categories={kubeops,appscode}
type LicenseProxyserverManager struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LicenseProxyserverManagerSpec `json:"spec,omitempty"`
}

// LicenseProxyserverManagerSpec is the schema for LicenseProxyserverManager Operator values file
type LicenseProxyserverManagerSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string `json:"fullnameOverride"`
	RegistryFQDN     string `json:"registryFQDN"`
	Image            string `json:"image"`
	// +optional
	Tag             string        `json:"tag"`
	ImagePullPolicy string        `json:"imagePullPolicy"`
	Placement       PlacementSpec `json:"placement"`
	Platform        PlatformSpec  `json:"platform"`
	//+optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
	// List of sources to populate environment variables in the container.
	// The keys defined within a source must be a C_IDENTIFIER. All invalid keys
	// will be reported as an event when the container is starting. When a key exists in multiple
	// sources, the value associated with the last source will take precedence.
	// Values defined by an Env with a duplicate key will take precedence.
	// Cannot be updated.
	// +optional
	// +listType=atomic
	EnvFrom []core.EnvFromSource `json:"envFrom"`
	// List of environment variables to set in the container.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=name
	Env []core.EnvVar `json:"env"`
	// +optional
	Distro shared.DistroSpec `json:"distro"`
}

type PlacementSpec struct {
	Create bool   `json:"create"`
	Name   string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LicenseProxyserverManagerList is a list of LicenseProxyserverManagers
type LicenseProxyserverManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LicenseProxyserverManager CRD objects
	Items []LicenseProxyserverManager `json:"items,omitempty"`
}
