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
	ResourceKindIngressNginx = "IngressNginx"
	ResourceIngressNginx     = "ingressnginx"
	ResourceIngressNginxs    = "ingressnginxs"
)

// IngressNginx defines the schama for IngressNginx Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=ingressnginxs,singular=ingressnginx,categories={kubeops,appscode}
type IngressNginx struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IngressNginxSpec `json:"spec,omitempty"`
}

// IngressNginxSpec is the schema for IngressNginx Operator values file
type IngressNginxSpec struct {
	Controller IngressNginxController `json:"controller"`
	// +optional
	TCP map[string]string `json:"tcp,omitempty"`
}

type IngressNginxController struct {
	Image                IngressNginxControllerImage                `json:"image"`
	Config               map[string]string                          `json:"config"`
	HostPort             *IngressNginxControllerHostPort            `json:"hostPort,omitempty"`
	IngressClassByName   bool                                       `json:"ingressClassByName"`
	IngressClassResource IngressNginxControllerIngressClassResource `json:"ingressClassResource"`
	Kind                 string                                     `json:"kind,omitempty"`
	//+optional
	NodeSelector map[string]string              `json:"nodeSelector"`
	Service      *IngressNginxControllerService `json:"service,omitempty"`
	//+optional
	Resources         core.ResourceRequirements     `json:"resources"`
	AdmissionWebhooks IngressNginxAdmissionWebhooks `json:"admissionWebhooks"`
}

type IngressNginxControllerImage struct {
	Chroot                   bool   `json:"chroot"`
	Registry                 string `json:"registry"`
	Image                    string `json:"image"`
	Tag                      string `json:"tag"`
	Digest                   string `json:"digest"`
	DigestChroot             string `json:"digestChroot"`
	PullPolicy               string `json:"pullPolicy"`
	RunAsUser                int    `json:"runAsUser"`
	AllowPrivilegeEscalation bool   `json:"allowPrivilegeEscalation"`
}

type IngressNginxControllerHostPort struct {
	Enabled bool `json:"enabled"`
}

type IngressNginxControllerIngressClassResource struct {
	ControllerValue string `json:"controllerValue"`
	Enabled         bool   `json:"enabled"`
	Name            string `json:"name"`
}

type IngressNginxControllerService struct {
	External IngressNginxControllerServiceExternal `json:"external"`
}

type IngressNginxControllerServiceExternal struct {
	Enabled bool `json:"enabled"`
}

type IngressNginxAdmissionWebhooks struct {
	Patch IngressNginxAdmissionWebhooksPatch `json:"patch"`
}

type IngressNginxAdmissionWebhooksPatch struct {
	Image IngressNginxAdmissionWebhooksImage `json:"image"`
}

type IngressNginxAdmissionWebhooksImage struct {
	Registry   string `json:"registry"`
	Image      string `json:"image"`
	Tag        string `json:"tag"`
	Digest     string `json:"digest"`
	PullPolicy string `json:"pullPolicy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IngressNginxList is a list of IngressNginxs
type IngressNginxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IngressNginx CRD objects
	Items []IngressNginx `json:"items,omitempty"`
}
