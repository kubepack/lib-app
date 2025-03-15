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
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	ResourceKindOpenfga = "Openfga"
	ResourceOpenfga     = "openfga"
	ResourceOpenfgas    = "openfgas"
)

// Openfga defines the schama for Openfga Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=openfgas,singular=openfga,categories={kubeops,appscode}
type Openfga struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpenfgaSpec `json:"spec,omitempty"`
}

type OpenfgaImageReference struct {
	Repository string `json:"repository"`
	PullPolicy string `json:"pullPolicy"`
	Tag        string `json:"tag"`
}

type OpenfgaServiceAccountSpec struct {
	Create      bool              `json:"create"`
	Annotations map[string]string `json:"annotations"`
	Name        string            `json:"name"`
}

type OpenfgaServiceSpec struct {
	Annotations map[string]string `json:"annotations"`
	Type        string            `json:"type"`
	Port        int               `json:"port"`
}

// OpenfgaSpec is the schema for Openfga Operator values file
type OpenfgaSpec struct {
	ReplicaCount                              int                         `json:"replicaCount"`
	Image                                     OpenfgaImageReference       `json:"image"`
	ImagePullSecrets                          []core.LocalObjectReference `json:"imagePullSecrets"`
	NameOverride                              string                      `json:"nameOverride"`
	FullnameOverride                          string                      `json:"fullnameOverride"`
	CommonLabels                              map[string]string           `json:"commonLabels"`
	ServiceAccount                            OpenfgaServiceAccountSpec   `json:"serviceAccount"`
	Annotations                               map[string]string           `json:"annotations"`
	PodAnnotations                            map[string]string           `json:"podAnnotations"`
	PodExtraLabels                            map[string]string           `json:"podExtraLabels"`
	ExtraEnvVars                              []core.EnvVar               `json:"extraEnvVars"`
	ExtraVolumes                              []core.Volume               `json:"extraVolumes"`
	ExtraVolumeMounts                         []core.VolumeMount          `json:"extraVolumeMounts"`
	ExtraInitContainers                       []core.Container            `json:"extraInitContainers"`
	PodSecurityContext                        core.PodSecurityContext     `json:"podSecurityContext"`
	SecurityContext                           core.SecurityContext        `json:"securityContext"`
	InitContainer                             OpenfgaImageReference       `json:"initContainer"`
	LivenessProbe                             OpenfgaProbe                `json:"livenessProbe"`
	ReadinessProbe                            OpenfgaProbe                `json:"readinessProbe"`
	StartupProbe                              OpenfgaProbe                `json:"startupProbe"`
	CustomLivenessProbe                       core.Probe                  `json:"customLivenessProbe"`
	CustomReadinessProbe                      core.Probe                  `json:"customReadinessProbe"`
	CustomStartupProbe                        core.Probe                  `json:"customStartupProbe"`
	Service                                   OpenfgaServiceSpec          `json:"service"`
	Telemetry                                 OpenfgaTelemetry            `json:"telemetry"`
	Datastore                                 OpenfgaDatastore            `json:"datastore"`
	Postgresql                                OpenfgaPostgresql           `json:"postgresql"`
	Mysql                                     OpenfgaMysql                `json:"mysql"`
	Grpc                                      OpenfgaGrpc                 `json:"grpc"`
	Http                                      OpenfgaHttp                 `json:"http"`
	Authn                                     OpenfgaAuthn                `json:"authn"`
	Playground                                OpenfgaPlayground           `json:"playground"`
	Profiler                                  OpenfgaProfiler             `json:"profiler"`
	Log                                       OpenfgaLog                  `json:"log"`
	CheckQueryCache                           OpenfgaCheckQueryCache      `json:"checkQueryCache"`
	Experimentals                             []string                    `json:"experimentals"`
	MaxTuplesPerWrite                         *int                        `json:"maxTuplesPerWrite"`
	MaxTypesPerAuthorizationModel             *int                        `json:"maxTypesPerAuthorizationModel"`
	MaxAuthorizationModelSizeInBytes          *int                        `json:"maxAuthorizationModelSizeInBytes"`
	MaxConcurrentReadsForCheck                *int                        `json:"maxConcurrentReadsForCheck"`
	MaxConcurrentReadsForListObjects          *int                        `json:"maxConcurrentReadsForListObjects"`
	ChangelogHorizonOffset                    *string                     `json:"changelogHorizonOffset"`
	ResolveNodeLimit                          *int                        `json:"resolveNodeLimit"`
	ResolveNodeBreadthLimit                   *int                        `json:"resolveNodeBreadthLimit"`
	ListObjectsDeadline                       *string                     `json:"listObjectsDeadline"`
	ListObjectsMaxResults                     *int                        `json:"listObjectsMaxResults"`
	ListUsersDeadline                         *string                     `json:"listUsersDeadline"`
	ListUsersMaxResults                       *int                        `json:"listUsersMaxResults"`
	MaxConcurrentReadsForListUsers            *int                        `json:"maxConcurrentReadsForListUsers"`
	RequestDurationDatastoreQueryCountBuckets []int                       `json:"requestDurationDatastoreQueryCountBuckets"`
	AllowWriting10Models                      *string                     `json:"allowWriting1_0Models"`
	AllowEvaluating10Models                   *string                     `json:"allowEvaluating1_0Models"`
	Ingress                                   AppIngress                  `json:"ingress"`
	Resources                                 core.ResourceRequirements   `json:"resources"`
	Autoscaling                               AutoscalingSpec             `json:"autoscaling"`
	NodeSelector                              map[string]string           `json:"nodeSelector"`
	Tolerations                               []core.Toleration           `json:"tolerations"`
	Affinity                                  core.Affinity               `json:"affinity"`
	Sidecars                                  []core.Container            `json:"sidecars"`
	Migrate                                   OpenfgaMigrate              `json:"migrate"`
	ExtraObjects                              []runtime.RawExtension      `json:"extraObjects"`
}
type OpenfgaProbe struct {
	Enabled    bool `json:"enabled"`
	core.Probe `json:",inline,omitempty"`
}

type OpenfgaTelemetry struct {
	Trace   OpenfgaTrace   `json:"trace"`
	Metrics OpenfgaMetrics `json:"metrics"`
}

type OpenfgaMetrics struct {
	Enabled             bool                  `json:"enabled"`
	ServiceMonitor      OpenfgaServiceMonitor `json:"serviceMonitor"`
	Addr                string                `json:"addr"`
	EnableRPCHistograms *bool                 `json:"enableRPCHistograms"`
	PodAnnotations      map[string]string     `json:"podAnnotations"`
}

type OpenfgaServiceMonitor struct {
	Enabled           bool                   `json:"enabled"`
	AdditionalLabels  map[string]string      `json:"additionalLabels"`
	Annotations       map[string]string      `json:"annotations"`
	JobLabel          string                 `json:"jobLabel"`
	Namespace         string                 `json:"namespace"`
	NamespaceSelector map[string]string      `json:"namespaceSelector"`
	ScrapeInterval    string                 `json:"scrapeInterval"`
	ScrapeTimeout     string                 `json:"scrapeTimeout"`
	TargetLabels      []runtime.RawExtension `json:"targetLabels"`
	Relabelings       []runtime.RawExtension `json:"relabelings"`
	MetricRelabelings []runtime.RawExtension `json:"metricRelabelings"`
}

type OpenfgaOltp struct {
	Endpoint *string        `json:"endpoint"`
	Tls      OpenfgaOltpTls `json:"tls"`
}

type OpenfgaOltpTls struct {
	Enabled bool `json:"enabled"`
}

type OpenfgaTrace struct {
	Enabled     bool         `json:"enabled"`
	Otlp        OpenfgaOltp  `json:"otlp"`
	SampleRatio *json.Number `json:"sampleRatio"`
}

type OpenfgaDatastore struct {
	Engine            string            `json:"engine"`
	Uri               *string           `json:"uri"`
	UriSecret         *string           `json:"uriSecret"`
	MaxCacheSize      *string           `json:"maxCacheSize"`
	MaxOpenConns      *string           `json:"maxOpenConns"`
	MaxIdleConns      *string           `json:"maxIdleConns"`
	ConnMaxIdleTime   *string           `json:"connMaxIdleTime"`
	ConnMaxLifetime   *string           `json:"connMaxLifetime"`
	ApplyMigrations   bool              `json:"applyMigrations"`
	WaitForMigrations bool              `json:"waitForMigrations"`
	MigrationType     string            `json:"migrationType"`
	Migrations        OpenfgaMigrations `json:"migrations"`
}

type OpenfgaMigrations struct {
	Resources core.ResourceRequirements `json:"resources"`
	Image     OpenfgaImageReference     `json:"image"`
}

type OpenfgaPostgresql struct {
	Enabled bool `json:"enabled"`
}

type OpenfgaMysql struct {
	Enabled bool `json:"enabled"`
}

type OpenfgaGrpc struct {
	Addr string     `json:"addr"`
	Tls  OpenfgaTls `json:"tls"`
}

type OpenfgaHttp struct {
	Enabled            bool       `json:"enabled"`
	Addr               string     `json:"addr"`
	Tls                OpenfgaTls `json:"tls"`
	UpstreamTimeout    *string    `json:"upstreamTimeout"`
	CorsAllowedOrigins []string   `json:"corsAllowedOrigins"`
	CorsAllowedHeaders []string   `json:"corsAllowedHeaders"`
}

type OpenfgaAuthn struct {
	Method    *string          `json:"method"`
	Preshared OpenfgaPreshared `json:"preshared"`
	Oidc      OpenfgaOidc      `json:"oidc"`
}

type OpenfgaPreshared struct {
	Keys []string `json:"keys"`
}

type OpenfgaOidc struct {
	Audience *string `json:"audience"`
	Issuer   *string `json:"issuer"`
}

type OpenfgaTls struct {
	Enabled bool    `json:"enabled"`
	Cert    *string `json:"cert"`
	Key     *string `json:"key"`
}

type OpenfgaPlayground struct {
	Enabled bool `json:"enabled"`
	Port    int  `json:"port"`
}

type OpenfgaProfiler struct {
	Enabled bool   `json:"enabled"`
	Addr    string `json:"addr"`
}

type OpenfgaLog struct {
	Level           string `json:"level"`
	Format          string `json:"format"`
	TimestampFormat string `json:"timestampFormat"`
}

type OpenfgaCheckQueryCache struct {
	Enabled bool    `json:"enabled"`
	Limit   *string `json:"limit"`
	Ttl     *string `json:"ttl"`
}

type OpenfgaMigrate struct {
	ExtraVolumes            []core.Volume      `json:"extraVolumes"`
	ExtraVolumeMounts       []core.VolumeMount `json:"extraVolumeMounts"`
	Sidecars                []core.Container   `json:"sidecars"`
	Annotations             map[string]*string `json:"annotations"`
	Labels                  map[string]*string `json:"labels"`
	Timeout                 *string            `json:"timeout"`
	Hook                    AceHook            `json:"hook"`
	TTLSecondsAfterFinished int                `json:"ttlSecondsAfterFinished"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenfgaList is a list of Openfgas
type OpenfgaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Openfga CRD objects
	Items []Openfga `json:"items,omitempty"`
}
