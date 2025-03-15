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
	catgwapi "go.bytebuilders.dev/catalog/api/gateway/v1alpha1"
	wizardsapi "go.bytebuilders.dev/ui-wizards/apis/wizards/v1alpha1"

	openviz_installer "go.openviz.dev/installer/apis/installer/v1alpha1"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	dnsapi "kubeops.dev/external-dns-operator/apis/external/v1alpha1"
)

const (
	ResourceKindAce = "Ace"
	ResourceAce     = "ace"
	ResourceAces    = "aces"
)

// Ace defines the schama for Ace Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=aces,singular=ace,categories={kubeops,appscode}
type Ace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AceSpec `json:"spec,omitempty"`
}

// AceSpec is the schema for Ace Operator values file
type AceSpec struct {
	PlatformUi   AcePlatformUi   `json:"platform-ui"`
	ClusterUi    AceClusterUi    `json:"cluster-ui"`
	Grafana      AceGrafana      `json:"grafana"`
	KubedbUi     AceKubedbUi     `json:"kubedb-ui"`
	PlatformApi  AcePlatformApi  `json:"platform-api"`
	IngressNginx AceIngressNginx `json:"ingress-nginx"`
	IngressDns   AceIngressDns   `json:"ingress-dns"`
	Nats         AceNats         `json:"nats"`
	NatsDns      AceNatsDns      `json:"nats-dns"`
	Trickster    AceTrickster    `json:"trickster"`
	Openfga      AceOpenfga      `json:"openfga"`
	S3proxy      AceS3proxy      `json:"s3proxy"`
	// KubeBindServer AceKubeBindServer `json:"kube-bind-server"`
	Global             AceGlobalValues           `json:"global"`
	Settings           Settings                  `json:"settings"`
	Image              ImageReference            `json:"image"`
	Kubectl            ImageReference            `json:"kubectl"`
	PodAnnotations     map[string]string         `json:"podAnnotations"`
	PodSecurityContext *core.PodSecurityContext  `json:"podSecurityContext"`
	SecurityContext    *core.SecurityContext     `json:"securityContext"`
	Resources          core.ResourceRequirements `json:"resources"`
	//+optional
	NodeSelector map[string]string                `json:"nodeSelector"`
	Tolerations  []core.Toleration                `json:"tolerations"`
	Affinity     *core.Affinity                   `json:"affinity"`
	Branding     AceBrandingSpec                  `json:"branding"`
	SetupJob     AceSetupJob                      `json:"setupJob"`
	ExtraObjects map[string]*runtime.RawExtension `json:"extraObjects"`
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
}

type AcePlatformUi struct {
	Enabled         bool `json:"enabled"`
	*PlatformUiSpec `json:",inline,omitempty"`
}

type AceAccountsUi struct {
	Enabled         bool `json:"enabled"`
	*AccountsUiSpec `json:",inline,omitempty"`
}

type AceClusterUi struct {
	Enabled        bool `json:"enabled"`
	*ClusterUiSpec `json:",inline,omitempty"`
}

type AceGrafana struct {
	Enabled      bool `json:"enabled"`
	*GrafanaSpec `json:",inline,omitempty"`
}

type AceKubedbUi struct {
	Enabled       bool `json:"enabled"`
	*KubedbUiSpec `json:",inline,omitempty"`
}

type AcePlatformApi struct {
	Enabled          bool `json:"enabled"`
	*PlatformApiSpec `json:",inline,omitempty"`
}

type AceIngressNginx struct {
	Enabled           bool `json:"enabled"`
	*IngressNginxSpec `json:",inline,omitempty"`
}

type AceIngressDns struct {
	Enabled bool                    `json:"enabled"`
	Spec    *dnsapi.ExternalDNSSpec `json:"spec,omitempty"`
}

type AceNats struct {
	Enabled   bool `json:"enabled"`
	*NatsSpec `json:",inline,omitempty"`
}

type AceNatsDns struct {
	Enabled bool                    `json:"enabled"`
	Spec    *dnsapi.ExternalDNSSpec `json:"spec,omitempty"`
}

type AceReloader struct {
	Enabled       bool `json:"enabled"`
	*ReloaderSpec `json:",inline,omitempty"`
}

type AceTrickster struct {
	Enabled                          bool `json:"enabled"`
	*openviz_installer.TricksterSpec `json:",inline,omitempty"`
}

type AceOpenfga struct {
	Enabled      bool   `json:"enabled"`
	DatastoreURI string `json:"datastoreURI"`
	*OpenfgaSpec `json:",inline,omitempty"`
}

type AceS3proxy struct {
	Enabled      bool `json:"enabled"`
	*S3proxySpec `json:",inline,omitempty"`
}

type AceGlobalValues struct {
	NameOverride     string                 `json:"nameOverride"`
	FullnameOverride string                 `json:"fullnameOverride"`
	Platform         AcePlatformSettings    `json:"platform"`
	License          string                 `json:"license"`
	Registry         string                 `json:"registry"`
	RegistryFQDN     string                 `json:"registryFQDN"`
	ImagePullSecrets []string               `json:"imagePullSecrets"`
	ServiceAccount   NatsServiceAccountSpec `json:"serviceAccount"`
	Monitoring       GlobalMonitoring       `json:"monitoring"`
	Infra            PlatformInfra          `json:"infra"`
}

type AcePlatformSettings struct {
	HostInfo           `json:",inline"`
	DeploymentType     DeploymentType `json:"deploymentType"`
	ProxyServiceDomain string         `json:"proxyServiceDomain,omitempty"`
	Token              string         `json:"token,omitempty"`
	OwnerID            int64          `json:"ownerID"`
	OwnerName          string         `json:"ownerName"`
}

type HostInfo struct {
	Host     string            `json:"host"`
	HostType catgwapi.HostType `json:"hostType"`
}

type GlobalMonitoring struct {
	Agent          string                   `json:"agent"`
	ServiceMonitor GlobalServiceMonitor     `json:"serviceMonitor"`
	Exporter       GlobalPrometheusExporter `json:"exporter"`
}

type GlobalServiceMonitor struct {
	Labels map[string]string `json:"labels"`
}

type GlobalPrometheusExporter struct {
	Resources core.ResourceRequirements `json:"resources"`
}

type PlatformInfra struct {
	StorageClass LocalObjectReference `json:"storageClass"`
	TLS          catgwapi.InfraTLS    `json:"tls"`
	DNS          InfraDns             `json:"dns"`
	Objstore     InfraObjstore        `json:"objstore"`
	Kubestash    KubeStashSpec        `json:"kubestash,omitempty"`
	Kms          InfraKms             `json:"kms"`
	Kubepack     InfraKubepack        `json:"kubepack"`
	Badger       InfraBadger          `json:"badger"`
	Invoice      InfraInvoice         `json:"invoice"`
	Fileserver   InfraFileserver      `json:"fileserver"`
}

type KubeStashSpec struct {
	// Schedule specifies the schedule for invoking backup sessions
	// +optional
	Schedule         string                      `json:"schedule,omitempty"`
	StorageRef       ObjectReference             `json:"storageRef"`
	RetentionPolicy  ObjectReference             `json:"retentionPolicy"`
	EncryptionSecret ObjectReference             `json:"encryptionSecret"`
	StorageSecret    wizardsapi.OptionalResource `json:"storageSecret"`
}

type InfraDns struct {
	catgwapi.GatewayDns `json:",inline,omitempty"`
	// +optional
	TargetIPs []string `json:"targetIPs"`
}

// +kubebuilder:validation:Enum=gcs;s3;azure;swift
type ObjstoreProvider string

const (
	ObjstoreProviderS3    ObjstoreProvider = "s3"
	ObjstoreProviderAzure ObjstoreProvider = "azure"
	ObjstoreProviderGCS   ObjstoreProvider = "gcs"
	ObjstoreProviderSwift ObjstoreProvider = "swift"
)

type InfraObjstore struct {
	Provider  ObjstoreProvider      `json:"provider"`
	Bucket    string                `json:"bucket"`
	Prefix    string                `json:"prefix,omitempty"`
	Endpoint  string                `json:"endpoint,omitempty"`
	Region    string                `json:"region,omitempty"`
	MountPath string                `json:"mountPath"`
	S3        *wizardsapi.S3Auth    `json:"s3,omitempty"`
	Azure     *wizardsapi.AzureAuth `json:"azure,omitempty"`
	GCS       *wizardsapi.GCSAuth   `json:"gcs,omitempty"`
}

type InfraKms struct {
	MasterKeyURL string `json:"masterKeyURL"`
}

type InfraKubepack struct {
	Bucket string `json:"bucket"`
	Prefix string `json:"prefix"`
}

type InfraBadger struct {
	MountPath string `json:"mountPath"`
	Levels    int    `json:"levels"`
}

type InfraInvoice struct {
	MountPath    string `json:"mountPath"`
	Bucket       string `json:"bucket"`
	Prefix       string `json:"prefix"`
	TrackerEmail string `json:"trackerEmail"`
}

type InfraFileserver struct {
	BaseURL string `json:"baseURL"`
	Bucket  string `json:"bucket"`
	Prefix  string `json:"prefix"`
}

type Settings struct {
	DB          DBSettings          `json:"db"`
	Cache       CacheSettings       `json:"cache"`
	Smtp        SmtpSettings        `json:"smtp"`
	Nats        NatsSettings        `json:"nats"`
	Platform    PlatformSettings    `json:"platform"`
	Security    SecuritySettings    `json:"security"`
	Grafana     GrafanaSettings     `json:"grafana"`
	InboxServer InboxServerSettings `json:"inboxServer"`
	Contract    ContractStorage     `json:"contract"`
	Firebase    FirebaseSettings    `json:"firebase"`
	// +optional
	Marketplace *MarketplaceSettings `json:"marketplace,omitempty"`
}

type DBSettings struct {
	Version        string                    `json:"version"`
	DatabaseName   string                    `json:"databaseName"`
	DeletionPolicy string                    `json:"deletionPolicy"`
	Persistence    PersistenceSpec           `json:"persistence"`
	Resources      core.ResourceRequirements `json:"resources"`
	Auth           BasicAuth                 `json:"auth"`
	LogSQL         bool                      `json:"logSQL"`
}

type CacheSettings struct {
	CacheInterval  int                       `json:"cacheInterval"`
	Version        string                    `json:"version"`
	DeletionPolicy string                    `json:"deletionPolicy"`
	Persistence    PersistenceSpec           `json:"persistence"`
	Resources      core.ResourceRequirements `json:"resources"`
	Auth           BasicAuth                 `json:"auth"`
}

type BasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SmtpSettings struct {
	Enabled         bool   `json:"enabled"`
	Host            string `json:"host"`
	TlsEnabled      bool   `json:"tlsEnabled"`
	From            string `json:"from"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	SubjectPrefix   string `json:"subjectPrefix"`
	SendAsPlainText bool   `json:"sendAsPlainText"`
}

type NatsSettings struct {
	ShardCount      int    `json:"shardCount"`
	Replics         int    `json:"replicas"`
	MountPath       string `json:"mountPath"`
	OperatorCreds   string `json:"operatorCreds"`
	OperatorJwt     string `json:"operatorJwt"`
	SystemCreds     string `json:"systemCreds"`
	SystemJwt       string `json:"systemJwt"`
	SystemPubKey    string `json:"systemPubKey"`
	SystemUserCreds string `json:"systemUserCreds"`
	AdminCreds      string `json:"adminCreds"`
	AdminUserCreds  string `json:"adminUserCreds"`
}

// +kubebuilder:validation:Enum=DEV;ALPHA;BETA;GA
type RunMode string

const (
	RunModeDev   RunMode = "DEV"
	RunModeAlpha RunMode = "ALPHA"
	RunModeBeta  RunMode = "BETA"
	RunModeGA    RunMode = "GA"
)

type PlatformSettings struct {
	AppName                      string  `json:"appName"`
	RunMode                      RunMode `json:"runMode"`
	ForcePrivate                 bool    `json:"forcePrivate"`
	DisableHttpGit               bool    `json:"disableHttpGit"`
	InstallLock                  bool    `json:"installLock"`
	RepositoryUploadEnabled      bool    `json:"repositoryUploadEnabled"`
	RepositoryUploadAllowedTypes *string `json:"repositoryUploadAllowedTypes"`
	RepositoryUploadMaxFileSize  int     `json:"repositoryUploadMaxFileSize"`
	RepositoryUploadMaxFiles     int     `json:"repositoryUploadMaxFiles"`
	ServiceEnableCaptcha         bool    `json:"serviceEnableCaptcha"`
	ServiceRegisterEmailConfirm  bool    `json:"serviceRegisterEmailConfirm"`
	ServiceDisableRegistration   bool    `json:"serviceDisableRegistration"`
	ServiceRequireSignInView     bool    `json:"serviceRequireSignInView"`
	ServiceEnableNotifyMail      bool    `json:"serviceEnableNotifyMail"`
	// +optional
	ServiceDomainWhiteList          []string `json:"serviceDomainWhiteList"`
	CookieName                      string   `json:"cookieName"`
	CookieRememberName              string   `json:"cookieRememberName"`
	CookieUsername                  string   `json:"cookieUsername"`
	ServerLandingPage               string   `json:"serverLandingPage"`
	LogMode                         string   `json:"logMode"`
	LogLevel                        string   `json:"logLevel"`
	OtherShowFooterBranding         bool     `json:"otherShowFooterBranding"`
	OtherShowFooterVersion          bool     `json:"otherShowFooterVersion"`
	OtherShowFooterTemplateLoadTime bool     `json:"otherShowFooterTemplateLoadTime"`
	EnableCSRFCookieHttpOnly        bool     `json:"enableCSRFCookieHttpOnly"`
	// +optional
	LoginURL string `json:"loginURL"`
	// +optional
	LogoutURL string `json:"logoutURL"`
}

type SecuritySettings struct {
	Oauth2JWTSecret string `json:"oauth2JWTSecret"`
	CsrfSecretKey   string `json:"csrfSecretKey"`
}

type GrafanaSettings struct {
	AppMode string `json:"appMode"`
	// +optional
	SecretKey string `json:"secretKey"`
}

type InboxServerSettings struct {
	JmapURL            string `json:"jmapURL"`
	WebAdminURL        string `json:"webAdminURL"`
	AdminJWTPrivateKey string `json:"adminJWTPrivateKey"`
}

type ContractStorage struct {
	Bucket        string `json:"bucket"`
	Prefix        string `json:"prefix"`
	LicenseBucket string `json:"licenseBucket"`
}

type FirebaseSettings struct {
	Project  string `json:"project"`
	Database string `json:"database"`
}

type MarketplaceSettings struct {
	AlertEmails           []string `json:"alertEmails"`
	SpreadsheetID         string   `json:"spreadsheetID"`
	SpreadsheetCredential string   `json:"spreadsheetCredential"`
	// /data/marketplace-credentials
	SpreadsheetCredentialMountPath string `json:"spreadsheetCredentialMountPath"`

	Aws   *AceOptionsAwsMarketplace   `json:"aws,omitempty"`
	Azure *AceOptionsAzureMarketplace `json:"azure,omitempty"`
	Gcp   *AceOptionsGcpMarketplace   `json:"gcp,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AceList is a list of Aces
type AceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ace CRD objects
	Items []Ace `json:"items,omitempty"`
}
