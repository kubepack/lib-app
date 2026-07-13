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

import gwapi "sigs.k8s.io/gateway-api/apis/v1"

type AppGateway struct {
	Enabled      bool `json:"enabled"`
	*GatewaySpec `json:",inline,omitempty"`
	Hosts        []GatewayHost `json:"hosts"`
}

type GatewaySpec struct {
	Annotations map[string]string `json:"annotations"`
	TLS         *GatewayTLS       `json:"tls"`
}

type GatewayTLS struct {
	Enabled bool                  `json:"enabled"`
	Secret  *LocalObjectReference `json:"secret"`
}

type GatewayHost struct {
	Host  string      `json:"host"`
	Paths GatewayPath `json:"paths"`
}

type GatewayPath struct {
	Path     string              `json:"path"`
	PathType gwapi.PathMatchType `json:"pathType,omitempty"`
}
