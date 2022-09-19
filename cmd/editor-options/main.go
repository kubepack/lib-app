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

package main

import (
	"encoding/json"
	"log"
	"os"

	appapi "kubepack.dev/lib-app/api/v1alpha1"

	flag "github.com/spf13/pflag"
	"sigs.k8s.io/yaml"
)

var (
	url     = "https://bundles.byte.builders/ui/"
	name    = "mongodb-editor-options"
	version = "v0.1.0"
)

func main() {
	flag.StringVar(&url, "url", url, "Chart repo url")
	flag.StringVar(&name, "name", name, "Name of bundle")
	flag.StringVar(&version, "version", version, "Version of bundle")
	flag.Parse()

	opts := appapi.EditorParameters{
		ValuesFile:  "values.yaml",
		ValuesPatch: nil,
	}

	err := os.MkdirAll("artifacts/"+name, 0o755)
	if err != nil {
		log.Fatal(err)
	}

	{
		data, err := yaml.Marshal(opts)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile("artifacts/"+name+"/editor.yaml", data, 0o644)
		if err != nil {
			log.Fatal(err)
		}
	}

	{
		data, err := json.MarshalIndent(opts, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile("artifacts/"+name+"/editor.json", data, 0o644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
