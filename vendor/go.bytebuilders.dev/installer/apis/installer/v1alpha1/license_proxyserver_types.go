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
)

const (
	ResourceKindLicenseProxyserver = "LicenseProxyserver"
	ResourceLicenseProxyserver     = "licenseproxyserver"
	ResourceLicenseProxyservers    = "licenseproxyservers"
)

// LicenseProxyserver defines the schama for LicenseProxyserver Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=licenseproxyservers,singular=licenseproxyserver,categories={kubeops,appscode}
type LicenseProxyserver struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LicenseProxyserverSpec `json:"spec,omitempty"`
}

// LicenseProxyserverSpec is the schema for LicenseProxyserver Operator values file
type LicenseProxyserverSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string    `json:"fullnameOverride"`
	ReplicaCount     int32     `json:"replicaCount"`
	RegistryFQDN     string    `json:"registryFQDN"`
	Image            Container `json:"image"`
	ImagePullPolicy  string    `json:"imagePullPolicy"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	CriticalAddon bool `json:"criticalAddon"`
	//+optional
	LogLevel int32 `json:"logLevel"`
	//+optional
	Annotations map[string]string `json:"annotations"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector" protobuf:"bytes,12,rep,name=nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations" protobuf:"bytes,13,rep,name=tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity" protobuf:"bytes,14,opt,name=affinity"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	ServiceAccount     ServiceAccountSpec       `json:"serviceAccount"`
	Apiserver          WebHookSpec              `json:"apiserver"`
	Monitoring         Monitoring               `json:"monitoring"`
	Platform           PlatformSpec             `json:"platform"`
	// +optional
	Licenses map[string]string `json:"licenses"`
	// +optional
	EncodedLicenses map[string]string `json:"encodedLicenses"`
}

type PlatformSpec struct {
	BaseURL string `json:"baseURL"`
	Token   string `json:"token"`
}

type ImageRef struct {
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}

type Container struct {
	ImageRef `json:",inline"`
	// Compute Resources required by the sidecar container.
	// +optional
	Resources core.ResourceRequirements `json:"resources"`
	// Security options the pod should run with.
	// +optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
}

type ServiceAccountSpec struct {
	Create bool `json:"create"`
	//+optional
	Name *string `json:"name"`
	//+optional
	Annotations map[string]string `json:"annotations"`
}

type WebHookSpec struct {
	GroupPriorityMinimum       int32           `json:"groupPriorityMinimum"`
	VersionPriority            int32           `json:"versionPriority"`
	UseKubeapiserverFqdnForAks bool            `json:"useKubeapiserverFqdnForAks"`
	Healthcheck                HealthcheckSpec `json:"healthcheck"`
	ServingCerts               ServingCerts    `json:"servingCerts"`
}

type HealthcheckSpec struct {
	//+optional
	Enabled bool `json:"enabled"`
}

type ServingCerts struct {
	Generate bool `json:"generate"`
	//+optional
	CaCrt string `json:"caCrt"`
	//+optional
	ServerCrt string `json:"serverCrt"`
	//+optional
	ServerKey string `json:"serverKey"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LicenseProxyserverList is a list of LicenseProxyservers
type LicenseProxyserverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LicenseProxyserver CRD objects
	Items []LicenseProxyserver `json:"items,omitempty"`
}
