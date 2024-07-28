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

package v1alpha1

import (
	"strconv"

	"go.bytebuilders.dev/resource-model/apis/cloud"
	"go.bytebuilders.dev/resource-model/crds"

	"k8s.io/apimachinery/pkg/labels"
	"kmodules.xyz/client-go/apiextensions"
)

func (_ Credential) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(SchemeGroupVersion.WithResource(ResourceCredentials))
}

func (cred *Credential) ApplyLabels(resourceName, credType string, ownerID int64) {
	labelMap := map[string]string{
		cloud.LabelResourceName:      resourceName,
		cloud.LabelCredentialType:    credType,
		cloud.LabelCredentialOwnerID: strconv.FormatInt(ownerID, 10),
	}
	cred.ObjectMeta.SetLabels(labelMap)
}

func (_ Credential) FormatLabels(resourceName, credType string, ownerID int64) labels.Selector {
	labelMap := make(map[string]string)
	if resourceName != "" {
		labelMap[cloud.LabelResourceName] = resourceName
	}
	if credType != "" {
		labelMap[cloud.LabelCredentialType] = credType
	}
	if ownerID != 0 {
		labelMap[cloud.LabelCredentialOwnerID] = strconv.FormatInt(ownerID, 10)
	}

	return labels.SelectorFromSet(labelMap)
}
