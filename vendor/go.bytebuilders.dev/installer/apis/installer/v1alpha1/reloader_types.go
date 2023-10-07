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
	ResourceKindReloader = "Reloader"
	ResourceReloader     = "reloader"
	ResourceReloaders    = "reloaders"
)

// Reloader defines the schama for Reloader Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=reloaders,singular=reloader,categories={kubeops,appscode}
type Reloader struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReloaderSpec `json:"spec,omitempty"`
}

// ReloaderSpec is the schema for Reloader Operator values file
type ReloaderSpec struct {
	Global     ReloaderGlobal     `json:"global"`
	Kubernetes ReloaderKubernetes `json:"kubernetes"`
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string          `json:"fullnameOverride"`
	Reloader         ReloaderDetails `json:"reloader"`
}

type ReloaderGlobal struct {
	ImagePullSecrets []string `json:"imagePullSecrets"`
}

type ReloaderKubernetes struct {
	Host string `json:"host"`
}

type ReloaderDetails struct {
	IsArgoRollouts         bool                        `json:"isArgoRollouts"`
	IsOpenshift            bool                        `json:"isOpenshift"`
	IgnoreSecrets          bool                        `json:"ignoreSecrets"`
	EnableHA               bool                        `json:"enableHA"`
	IgnoreConfigMaps       bool                        `json:"ignoreConfigMaps"`
	SyncAfterRestart       bool                        `json:"syncAfterRestart"`
	ReloadOnCreate         bool                        `json:"reloadOnCreate"`
	ReloadStrategy         string                      `json:"reloadStrategy"`
	IgnoreNamespaces       string                      `json:"ignoreNamespaces"`
	NamespaceSelector      string                      `json:"namespaceSelector"`
	LogFormat              string                      `json:"logFormat"`
	WatchGlobally          bool                        `json:"watchGlobally"`
	ReadOnlyRootFileSystem bool                        `json:"readOnlyRootFileSystem"`
	Legacy                 ReloaderLegacy              `json:"legacy"`
	MatchLabels            map[string]string           `json:"matchLabels"`
	Deployment             ReloaderDeploymentSpec      `json:"deployment"`
	Service                ReloaderServiceSpec         `json:"service"`
	Rbac                   ReloaderRbacSpec            `json:"rbac"`
	ServiceAccount         ReloaderServiceAccountSpec  `json:"serviceAccount"`
	CustomAnnotations      map[string]string           `json:"custom_annotations"`
	ServiceMonitor         ReloaderServiceMonitorSpec  `json:"serviceMonitor"`
	PodMonitor             ReloaderPodMonitorSpec      `json:"podMonitor"`
	PodDisruptionBudget    ReloaderPodDisruptionBudget `json:"podDisruptionBudget"`
}

type ReloaderLegacy struct {
	Rbac bool `json:"rbac"`
}

type ReloaderDeploymentSpec struct {
	Replicas int `json:"replicas"`
	//+optional
	NodeSelector             map[string]string         `json:"nodeSelector"`
	Affinity                 *core.Affinity            `json:"affinity"`
	SecurityContext          *core.PodSecurityContext  `json:"securityContext"`
	ContainerSecurityContext *core.SecurityContext     `json:"containerSecurityContext"`
	Tolerations              []core.Toleration         `json:"tolerations"`
	Annotations              map[string]string         `json:"annotations"`
	Labels                   ReloaderLabels            `json:"labels"`
	Image                    ReloaderImageReference    `json:"image"`
	Env                      ReloaderEnvVars           `json:"env"`
	LivenessProbe            *core.Probe               `json:"livenessProbe"`
	ReadinessProbe           *core.Probe               `json:"readinessProbe"`
	Resources                core.ResourceRequirements `json:"resources"`
	Pod                      ReloaderPodSpec           `json:"pod"`
	PriorityClassName        string                    `json:"priorityClassName"`
}

type ReloaderLabels struct {
	Provider string `json:"provider"`
	Group    string `json:"group"`
	Version  string `json:"version"`
}

type ReloaderImageReference struct {
	Name       string `json:"name"`
	Tag        string `json:"tag"`
	PullPolicy string `json:"pullPolicy"`
}

type ReloaderEnvVars struct {
	Open     map[string]string `json:"open"`
	Secret   map[string]string `json:"secret"`
	Field    map[string]string `json:"field"`
	Existing *ReloaderExisting `json:"existing"`
}

type ReloaderExisting struct {
	ExistingSecretName map[string]string `json:"existing_secret_name,omitempty"`
}

type ReloaderPodSpec struct {
	Annotations map[string]string `json:"annotations"`
}

type ReloaderServiceSpec struct {
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
	Port        int               `json:"port,omitempty"`
}

type ReloaderRbacSpec struct {
	Enabled bool              `json:"enabled"`
	Labels  map[string]string `json:"labels"`
}

type ReloaderServiceAccountSpec struct {
	Create      bool              `json:"create"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	Name        *string           `json:"name"`
}

type ReloaderServiceMonitorSpec struct {
	Enabled     bool              `json:"enabled"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	HonorLabels bool              `json:"honorLabels"`
	// TODO(tamal): simplified, intentionally wrong
	MetricRelabelings []string `json:"metricRelabelings"`
	// TODO(tamal): simplified, intentionally wrong
	Relabelings  []string `json:"relabelings"`
	TargetLabels []string `json:"targetLabels"`
}

type ReloaderPodMonitorSpec struct {
	Enabled     bool              `json:"enabled"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	HonorLabels bool              `json:"honorLabels"`
	// TODO(tamal): simplified, intentionally wrong
	MetricRelabelings []string `json:"metricRelabelings"`
	// TODO(tamal): simplified, intentionally wrong
	Relabelings     []string `json:"relabelings"`
	PodTargetLabels []string `json:"podTargetLabels"`
}

type ReloaderPodDisruptionBudget struct {
	Enabled bool `json:"enabled"`
}

// EnvVar represents an environment variable present in a Container.
type EnvVar struct {
	// Name of the environment variable. Must be a C_IDENTIFIER.
	Name string `json:"name"`
	// +optional
	Value string `json:"value,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ReloaderList is a list of Reloaders
type ReloaderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Reloader CRD objects
	Items []Reloader `json:"items,omitempty"`
}
