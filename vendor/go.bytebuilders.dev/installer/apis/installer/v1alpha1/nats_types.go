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
	networking "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	ResourceKindNats = "Nats"
	ResourceNats     = "nats"
	ResourceNatss    = "natss"
)

// Nats defines the schama for Nats Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=natss,singular=nats,categories={kubeops,appscode}
type Nats struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NatsSpec `json:"spec,omitempty"`
}

// NatsSpec is the schema for Nats Operator values file
type NatsSpec struct {
	Nats                      NatsServerSpec                  `json:"nats"`
	Mqtt                      NatsMqttSpec                    `json:"mqtt"`
	NameOverride              string                          `json:"nameOverride"`
	NamespaceOverride         string                          `json:"namespaceOverride"`
	ImagePullSecrets          []string                        `json:"imagePullSecrets"`
	SecurityContext           *core.SecurityContext           `json:"securityContext"`
	Affinity                  *core.Affinity                  `json:"affinity"`
	PriorityClassName         *string                         `json:"priorityClassName"`
	TopologyKeys              []string                        `json:"topologyKeys"`
	TopologySpreadConstraints []core.TopologySpreadConstraint `json:"topologySpreadConstraints"`
	PodAnnotations            map[string]string               `json:"podAnnotations"`
	PodDisruptionBudget       NatsPodDisruptionBudgetSpec     `json:"podDisruptionBudget"`
	NodeSelector              map[string]string               `json:"nodeSelector"`
	Tolerations               []core.Toleration               `json:"tolerations"`
	StatefulSetAnnotations    map[string]string               `json:"statefulSetAnnotations"`
	StatefulSetPodLabels      map[string]string               `json:"statefulSetPodLabels"`
	ServiceAnnotations        map[string]string               `json:"serviceAnnotations"`
	AdditionalContainers      []core.Container                `json:"additionalContainers"`
	AdditionalVolumes         []core.Volume                   `json:"additionalVolumes"`
	AdditionalVolumeMounts    []core.VolumeMount              `json:"additionalVolumeMounts"`
	Cluster                   NatsClusterSpec                 `json:"cluster"`
	Leafnodes                 NatsLeafnodesSpec               `json:"leafnodes"`
	Gateway                   NatsGatewaySpec                 `json:"gateway"`
	Bootconfig                NatsBootconfigSpec              `json:"bootconfig"`
	Natsbox                   NatsboxSpec                     `json:"natsbox"`
	Reloader                  NatsReloaderSpec                `json:"reloader"`
	Exporter                  NatsExporterSpec                `json:"exporter"`
	Auth                      NatsAuthSpec                    `json:"auth"`
	Websocket                 NatsWebsocketSpec               `json:"websocket"`
	NetworkPolicy             NatsNetworkPolicySpec           `json:"networkPolicy"`
	K8SClusterDomain          string                          `json:"k8sClusterDomain"`
	UseFQDN                   bool                            `json:"useFQDN"`
	CommonLabels              map[string]string               `json:"commonLabels"`
	PodManagementPolicy       string                          `json:"podManagementPolicy"`
	PidVolume                 NatsTempVolume                  `json:"pidVolume"`
	AdvertiseconfigVolume     NatsTempVolume                  `json:"advertiseconfigVolume"`
}

type NatsServerSpec struct {
	Image                         NatsImageRef              `json:"image"`
	ServerNamePrefix              string                    `json:"serverNamePrefix"`
	ServerTags                    []string                  `json:"serverTags"`
	Profiling                     NatsServerProfilingSpec   `json:"profiling"`
	Healthcheck                   NatsServerHealthcheckSpec `json:"healthcheck"`
	HostNetwork                   bool                      `json:"hostNetwork"`
	DnsPolicy                     string                    `json:"dnsPolicy"`
	ConfigChecksumAnnotation      bool                      `json:"configChecksumAnnotation"`
	SecurityContext               *core.SecurityContext     `json:"securityContext"`
	ExternalAccess                bool                      `json:"externalAccess"`
	Advertise                     bool                      `json:"advertise"`
	ServiceAccount                *NatsServiceAccountSpec   `json:"serviceAccount"`
	ConnectRetries                int                       `json:"connectRetries"`
	SelectorLabels                map[string]string         `json:"selectorLabels"`
	Resources                     core.ResourceRequirements `json:"resources"`
	Client                        NatsServerClientSpec      `json:"client"`
	ExtraEnv                      []string                  `json:"extraEnv"`
	Limits                        NatsServerLimitsSpec      `json:"limits"`
	TerminationGracePeriodSeconds *int64                    `json:"terminationGracePeriodSeconds"`
	Logging                       NatsLoggingSpec           `json:"logging"`
	Mappings                      NatsMappings              `json:"mappings"`
	Jetstream                     JetstreamSpec             `json:"jetstream"`
	TLS                           *NatsServerTLSSpec        `json:"tls,omitempty"`
}

type NatsServerProfilingSpec struct {
	Enabled bool `json:"enabled"`
	Port    int  `json:"port"`
}

type NatsServerHealthcheckSpec struct {
	DetectHealthz                  bool           `json:"detectHealthz"`
	EnableHealthz                  bool           `json:"enableHealthz"`
	EnableHealthzLivenessReadiness bool           `json:"enableHealthzLivenessReadiness"`
	Liveness                       *LivenessProbe `json:"liveness"`
	Readiness                      *Probe         `json:"readiness"`
	Startup                        *Probe         `json:"startup"`
}

type LivenessProbe struct {
	Probe                         `json:",inline"`
	TerminationGracePeriodSeconds *int64 `json:"terminationGracePeriodSeconds,omitempty"`
}

type Probe struct {
	Enabled             bool `json:"enabled"`
	InitialDelaySeconds int  `json:"initialDelaySeconds"`
	TimeoutSeconds      int  `json:"timeoutSeconds"`
	PeriodSeconds       int  `json:"periodSeconds"`
	SuccessThreshold    int  `json:"successThreshold"`
	FailureThreshold    int  `json:"failureThreshold"`
}

type NatsServiceAccountSpec struct {
	Create      bool              `json:"create"`
	Annotations map[string]string `json:"annotations"`
	Name        string            `json:"name"`
}

type NatsServerClientSpec struct {
	Port     int    `json:"port"`
	PortName string `json:"portName"`
}

type NatsServerLimitsSpec struct {
	MaxConnections      *string `json:"maxConnections"`
	MaxSubscriptions    *string `json:"maxSubscriptions"`
	MaxControlLine      *string `json:"maxControlLine"`
	MaxPayload          *string `json:"maxPayload"`
	WriteDeadline       *string `json:"writeDeadline"`
	MaxPending          *string `json:"maxPending"`
	MaxPings            *string `json:"maxPings"`
	PingInterval        *string `json:"pingInterval"`
	LameDuckGracePeriod string  `json:"lameDuckGracePeriod"`
	LameDuckDuration    string  `json:"lameDuckDuration"`
}

type NatsLoggingSpec struct {
	Debug                 *bool   `json:"debug"`
	Trace                 *bool   `json:"trace"`
	Logtime               *string `json:"logtime"`
	ConnectErrorReports   *string `json:"connectErrorReports"`
	ReconnectErrorReports *string `json:"reconnectErrorReports"`
}

type NatsMappings struct{}

type JetstreamSpec struct {
	Enabled               bool                 `json:"enabled"`
	Domain                *string              `json:"domain"`
	UniqueTag             *string              `json:"uniqueTag"`
	MaxOutstandingCatchup *string              `json:"max_outstanding_catchup"`
	Encryption            runtime.RawExtension `json:"encryption"`
	MemStorage            JetstreamMemStorage  `json:"memStorage"`
	FileStorage           JetstreamFileStorage `json:"fileStorage"`
}

type JetstreamMemStorage struct {
	Enabled bool              `json:"enabled"`
	Size    resource.Quantity `json:"size"`
}

type JetstreamFileStorage struct {
	Enabled          bool   `json:"enabled"`
	StorageDirectory string `json:"storageDirectory"`
	// +optional
	StorageClassName string            `json:"storageClassName,omitempty"`
	Size             resource.Quantity `json:"size"`
	AccessModes      []string          `json:"accessModes"`
	Annotations      map[string]string `json:"annotations"`
}

type NatsServerTLSSpec struct {
	AllowNonTLS bool                 `json:"allowNonTLS"`
	Secret      LocalObjectReference `json:"secret"`
	Ca          string               `json:"ca"`
	Cert        string               `json:"cert"`
	Key         string               `json:"key"`
}

type NatsMqttSpec struct {
	Enabled       bool     `json:"enabled"`
	AckWait       string   `json:"ackWait"`
	MaxAckPending int      `json:"maxAckPending"`
	TLS           *TLSSpec `json:"tls,omitempty"`
}

type NatsPodDisruptionBudgetSpec struct {
	Enabled        bool `json:"enabled"`
	MaxUnavailable int  `json:"maxUnavailable"`
	MinAvailable   int  `json:"minAvailable,omitempty"`
}

type NatsClusterSpec struct {
	Enabled       bool               `json:"enabled"`
	Replicas      int                `json:"replicas"`
	NoAdvertise   bool               `json:"noAdvertise"`
	ExtraRoutes   []string           `json:"extraRoutes"`
	Authorization *NatsAuthorization `json:"authorization,omitempty"`
	TLS           *TLSSpec           `json:"tls,omitempty"`
}

type NatsLeafnodesSpec struct {
	Enabled     bool                        `json:"enabled"`
	Port        int                         `json:"port,omitempty"`
	NoAdvertise bool                        `json:"noAdvertise"`
	Remotes     []NatsLeafnodeRemoteAddress `json:"remotes,omitempty"`
	TLS         *TLSSpec                    `json:"tls,omitempty"`
}

type NatsLeafnodeRemoteAddress struct {
	URL string `json:"url"`
}

type NatsGatewaySpec struct {
	Enabled              bool                 `json:"enabled"`
	Port                 int                  `json:"port,omitempty"`
	Name                 string               `json:"name"`
	Authorization        *NatsAuthorization   `json:"authorization,omitempty"`
	RejectUnknownCluster bool                 `json:"rejectUnknownCluster,omitempty"`
	Advertise            string               `json:"advertise,omitempty"`
	Gateways             []NatsGatewayAddress `json:"gateways,omitempty"`
	TLS                  *TLSSpec             `json:"tls,omitempty"`
}

type NatsAuthorization struct {
	User     string  `json:"user"`
	Password string  `json:"password"`
	Timeout  float64 `json:"timeout"`
}

type NatsGatewayAddress struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type TLSSpec struct {
	Secret LocalObjectReference `json:"secret"`
	Ca     string               `json:"ca"`
	Cert   string               `json:"cert"`
	Key    string               `json:"key"`
}

type NatsImageRef struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	PullPolicy string `json:"pullPolicy"`
	Registry   string `json:"registry,omitempty"`
}

type NatsBootconfigSpec struct {
	Image           NatsImageRef          `json:"image"`
	SecurityContext *core.SecurityContext `json:"securityContext"`
}

type NatsboxSpec struct {
	Enabled          bool                  `json:"enabled"`
	Image            NatsImageRef          `json:"image"`
	SecurityContext  *core.SecurityContext `json:"securityContext"`
	AdditionalLabels map[string]string     `json:"additionalLabels"`
	ImagePullSecrets []string              `json:"imagePullSecrets"`
	PodAnnotations   map[string]string     `json:"podAnnotations"`
	PodLabels        map[string]string     `json:"podLabels"`
	Affinity         *core.Affinity        `json:"affinity"`
	//+optional
	NodeSelector      map[string]string  `json:"nodeSelector"`
	Tolerations       []core.Toleration  `json:"tolerations"`
	ExtraVolumeMounts []core.VolumeMount `json:"extraVolumeMounts"`
	ExtraVolumes      []core.Volume      `json:"extraVolumes"`
}

type NatsReloaderSpec struct {
	Enabled         bool                  `json:"enabled"`
	Image           NatsImageRef          `json:"image"`
	SecurityContext *core.SecurityContext `json:"securityContext"`
	ExtraConfigs    []string              `json:"extraConfigs"`
}

type NatsExporterSpec struct {
	Enabled         bool                           `json:"enabled"`
	Image           NatsImageRef                   `json:"image"`
	PortName        string                         `json:"portName"`
	SecurityContext *core.SecurityContext          `json:"securityContext"`
	Resources       core.ResourceRequirements      `json:"resources"`
	Args            []string                       `json:"args"`
	ServiceMonitor  NatsExporterServiceMonitorSpec `json:"serviceMonitor"`
}

type NatsExporterServiceMonitorSpec struct {
	Enabled       bool              `json:"enabled"`
	Namespace     string            `json:"namespace,omitempty"`
	Labels        map[string]string `json:"labels"`
	Annotations   map[string]string `json:"annotations"`
	Path          string            `json:"path,omitempty"`
	Interval      string            `json:"interval,omitempty"`
	ScrapeTimeout string            `json:"scrapeTimeout,omitempty"`
}

type NatsAuthSpec struct {
	Enabled       bool                 `json:"enabled"`
	Operatorjwt   *NatsOperatorJWTSpec `json:"operatorjwt,omitempty"`
	SystemAccount *string              `json:"systemAccount,omitempty"`
	Resolver      NatsResolverSpec     `json:"resolver"`
}

type NatsOperatorJWTSpec struct {
	ConfigMap ConfigMapKeySelector `json:"configMap"`
}

type ConfigMapKeySelector struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type NatsResolverSpec struct {
	Type          string                `json:"type"`
	AllowDelete   bool                  `json:"allowDelete"`
	Interval      string                `json:"interval"`
	Operator      *string               `json:"operator"`
	SystemAccount *string               `json:"systemAccount"`
	Store         NatsResolverStoreSpec `json:"store"`
	// +optional
	ResolverPreload map[string]string `json:"resolverPreload,omitempty"`
}

type NatsResolverStoreSpec struct {
	Dir  string            `json:"dir"`
	Size resource.Quantity `json:"size"`
	// +optional
	StorageClassName string `json:"storageClassName,omitempty"`
}

type NatsWebsocketSpec struct {
	Enabled        bool     `json:"enabled"`
	Port           int      `json:"port"`
	NoTLS          bool     `json:"noTLS"`
	SameOrigin     bool     `json:"sameOrigin"`
	AllowedOrigins []string `json:"allowedOrigins"`
	Advertise      string   `json:"advertise,omitempty"`
	// +optional
	HandshakeTimeout string   `json:"handshakeTimeout,omitempty"`
	TLS              *TLSSpec `json:"tls,omitempty"`
}

type NatsAppProtocolSpec struct {
	Enabled bool `json:"enabled"`
}

type NatsNetworkPolicySpec struct {
	Enabled                 bool                                  `json:"enabled"`
	AllowExternal           bool                                  `json:"allowExternal"`
	ExtraIngress            []networking.NetworkPolicyIngressRule `json:"extraIngress"`
	ExtraEgress             []networking.NetworkPolicyEgressRule  `json:"extraEgress"`
	IngressNSMatchLabels    map[string]string                     `json:"ingressNSMatchLabels"`
	IngressNSPodMatchLabels map[string]string                     `json:"ingressNSPodMatchLabels"`
}

type NatsTempVolume struct {
	EmptyDir EmptyDir `json:"emptyDir"`
}

type EmptyDir struct{}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NatsList is a list of Natss
type NatsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Nats CRD objects
	Items []Nats `json:"items,omitempty"`
}
