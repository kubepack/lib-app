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

package fusion

const resourceTemplate = `{{"{{- with .Values.resources }}"}}
{{"{{- with ."}}{{ .key }}{{" }}"}}
{{"{{- . | toYaml }}"}}
{{"{{- end }}"}}
{{"{{- end }}"}}
`

const valuesMetadataSchema = `properties:
  metadata:
    properties:
      release:
        properties:
          name:
            type: string
          namespace:
            type: string
        required:
        - name
        - namespace
        type: object
      resource:
        description: ResourceID identifies a resource
        properties:
          group:
            type: string
          kind:
            description: Kind is the serialized kind of the resource.  It is normally
              CamelCase and singular.
            type: string
          name:
            description: 'Name is the plural name of the resource to serve.  It must
              match the name of the CustomResourceDefinition-registration too: plural.group
              and it must be all lowercase.'
            type: string
          scope:
            description: ResourceScope is an enum defining the different scopes available
              to a custom resource
            type: string
          version:
            type: string
        required:
        - group
        - kind
        - name
        - scope
        - version
        type: object
    required:
    - release
    - resource
    type: object
type: object`
