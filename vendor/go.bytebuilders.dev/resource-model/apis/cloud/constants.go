/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloud

import (
	"go/build"
	"path/filepath"
)

const (
	LabelResourceName      = "byte.builders/resource-name"
	LabelCredentialType    = "byte.builders/cluster-credential-type"
	LabelCredentialOwnerID = "byte.builders/cluster-credential-owner-id"

	KeyCloudProvider = "cloud.appscode.com/provider"
)

const (
	GCE          string = "gce"
	DigitalOcean string = "digitalocean"
	Packet       string = "packet"
	AWS          string = "aws"
	Azure        string = "azure"
	AzureStorage string = "azureStorage"
	Vultr        string = "vultr"
	Linode       string = "linode"
	Scaleway     string = "scaleway"
)

var DataDir string

func init() {
	DataDir = filepath.Join(build.Default.GOPATH, "src/go.bytebuilders.dev/resource-model/data")
}
