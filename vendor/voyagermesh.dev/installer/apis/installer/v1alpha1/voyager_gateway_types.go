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
	ResourceKindVoyagerGateway = "VoyagerGateway"
	ResourceVoyagerGateway     = "oyagergateway"
	ResourceVoyagerGateways    = "oyagergateways"
)

// VoyagerGateway defines the schama for VoyagerGateway operator installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
type VoyagerGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VoyagerGatewaySpec `json:"spec,omitempty"`
}

// VoyagerGatewaySpec is the schema for Operator Operator values file
type VoyagerGatewaySpec struct {
	Global                  VoyagerGatewayGlobal    `json:"global"`
	PodDisruptionBudget     PodDisruptionBudgetSpec `json:"podDisruptionBudget"`
	Deployment              DeploymentSpec          `json:"deployment"`
	Config                  EnvoyGatewayConfig      `json:"config"`
	CreateNamespace         bool                    `json:"createNamespace"`
	KubernetesClusterDomain string                  `json:"kubernetesClusterDomain"`
	Certgen                 CertgenSpec             `json:"certgen"`
	GatewayConverter        VoyagerGatewayConverter `json:"gateway-converter"`
}

type VoyagerGatewayGlobal struct {
	Images Images `json:"images"`
}

type Images struct {
	EnvoyGateway ImageDetails `json:"envoyGateway"`
	Ratelimit    ImageDetails `json:"ratelimit"`
}

type ImageDetails struct {
	Image       string                      `json:"image"`
	PullPolicy  string                      `json:"pullPolicy"`
	PullSecrets []core.LocalObjectReference `json:"pullSecrets"`
}

type PodDisruptionBudgetSpec struct {
	MinAvailable int `json:"minAvailable"`
}

type DeploymentSpec struct {
	EnvoyGateway EnvoyGatewayDeployment `json:"envoyGateway"`
	Ports        []Port                 `json:"ports"`
	Replicas     int                    `json:"replicas"`
	Pod          PodTemplateSpec        `json:"pod"`
}

type EnvoyGatewayDeployment struct {
	Image            ImageSpec                   `json:"image"`
	ImagePullPolicy  string                      `json:"imagePullPolicy"`
	ImagePullSecrets []core.LocalObjectReference `json:"imagePullSecrets"`
	Resources        core.ResourceRequirements   `json:"resources"`
	SecurityContext  *core.SecurityContext       `json:"securityContext,omitempty"`
}

type ImageSpec struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}

type Port struct {
	Name       string `json:"name"`
	Port       int    `json:"port"`
	TargetPort int    `json:"targetPort"`
}

type PodTemplateSpec struct {
	Affinity                  *core.Affinity                  `json:"affinity"`
	Annotations               map[string]string               `json:"annotations"`
	Labels                    map[string]string               `json:"labels"`
	TopologySpreadConstraints []core.TopologySpreadConstraint `json:"topologySpreadConstraints"`
	Tolerations               []core.Toleration               `json:"tolerations"`
}

type EnvoyGatewayConfig struct {
	EnvoyGateway EnvoyGatewaySpec `json:"envoyGateway"`
}

type EnvoyGatewaySpec struct {
	Gateway  GatewayControllerSpec `json:"gateway"`
	Provider GatewayProviderSpec   `json:"provider"`
	Logging  LoggingSpec           `json:"logging"`
}

type GatewayControllerSpec struct {
	ControllerName string `json:"controllerName"`
}

type GatewayProviderSpec struct {
	Type string `json:"type"`
}

type LoggingSpec struct {
	Level LoggingLevel `json:"level"`
}

type LoggingLevel struct {
	Default string `json:"default"`
}

type CertgenSpec struct {
	Job  CertgenJobSpec      `json:"job"`
	Rbac CertgenRbacMetadata `json:"rbac"`
}

type CertgenJobSpec struct {
	Annotations             map[string]string         `json:"annotations"`
	Resources               core.ResourceRequirements `json:"resources"`
	TtlSecondsAfterFinished int                       `json:"ttlSecondsAfterFinished"`
	SecurityContext         *core.SecurityContext     `json:"securityContext,omitempty"`
}

type CertgenRbacMetadata struct {
	Annotations map[string]string `json:"annotations"`
	Labels      map[string]string `json:"labels"`
}

type VoyagerGatewayConverter struct {
	Enabled               bool `json:"enabled"`
	*GatewayConverterSpec `json:",inline,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VoyagerGatewayList is a list of VoyagerGateways
type VoyagerGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VoyagerGateway CRD objects
	Items []VoyagerGateway `json:"items,omitempty"`
}
