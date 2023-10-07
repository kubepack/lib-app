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
	"k8s.io/apimachinery/pkg/api/resource"
)

type ImageReference struct {
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	PullPolicy string `json:"pullPolicy"`
}

type AceServiceSpec struct {
	Type string `json:"type"`
	Port int    `json:"port"`
}

type AutoscalingSpec struct {
	Enabled     bool `json:"enabled"`
	MinReplicas int  `json:"minReplicas"`
	MaxReplicas int  `json:"maxReplicas"`
	// +optional
	TargetCPUUtilizationPercentage int `json:"targetCPUUtilizationPercentage,omitempty"`
	// +optional
	TargetMemoryUtilizationPercentage int `json:"targetMemoryUtilizationPercentage,omitempty"`
}

// +kubebuilder:validation:Enum=prometheus.io;prometheus.io/operator;prometheus.io/builtin
type MonitoringAgent string

type Monitoring struct {
	Agent          MonitoringAgent      `json:"agent"`
	ServiceMonitor ServiceMonitorLabels `json:"serviceMonitor"`
}

type CustomMonitoring struct {
	Agent          MonitoringAgent      `json:"agent"`
	Port           int                  `json:"port"`
	ServiceMonitor ServiceMonitorLabels `json:"serviceMonitor"`
}

type ServiceMonitorLabels struct {
	// +optional
	Labels map[string]string `json:"labels"`
}

type AceInfra struct {
	StorageClass LocalObjectReference `json:"storageClass"`
	Objstore     ProviderMount        `json:"objstore"`
	Badger       VolumeMount          `json:"badger"`
	Invoice      VolumeMount          `json:"invoice"`
}

type LocalObjectReference struct {
	Name string `json:"name"`
}

// ObjectReference contains enough information to let you inspect or modify the referred object.
type ObjectReference struct {
	// Namespace of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace"`
	// Name of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name"`
}

type ProviderMount struct {
	Provider  string `json:"provider"`
	MountPath string `json:"mountPath"`
}

type VolumeMount struct {
	MountPath string `json:"mountPath"`
}

type AceSettings struct {
	CAProviderClass       string `json:"caProviderClass"`
	AceSettingsSecretName `json:"secretName"`
}

type AceSettingsSecretName struct {
	PlatformConfig string `json:"platformConfig"`
	GrafanaConfig  string `json:"grafanaConfig"`
	Objstore       string `json:"objstore"`
	Nats           string `json:"nats"`
}

type PersistenceSpec struct {
	Size resource.Quantity `json:"size"`
}
