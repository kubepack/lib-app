# UI Editor APIs

## Fuse chart from sample dir

```console
$ go run cmd/fuse-chart/*.go \
  --sample-dir=/home/tamal/go/src/github.com/appscode/kubedb-samples/mongodb/standalone/prometheus.io/backupconfiguration/stash/tls/custom-auth/standalone \
  --chart-dir=/home/tamal/go/src/go.bytebuilders.dev/ui-wizards/charts \
  --sample-name=mongodb \
  --resource.group=kubedb.com \
  --resource.version=v1alpha2 \
  --resource.name=mongodbs

$ go run cmd/gen-simple-editor/main.go --all --skipExisting=false
```

## Demo UI Editors

- GET "/bundleview"

- POST "/bundleview/orders", v1alpha1.BundleView{}

  - API PATH CHANGED from /deploy/orders -> /bundleview/orders

- GET "/packageview"

`http://localhost:4000/packageview?url=https://bundles.byte.builders/ui/&name=mongodb-editor-options&version=v0.1.0`

- POST "/packageview/orders"

`curl -X POST -H "Content-Type: application/json" -d @./artifacts/mongodb-editor/packageview_chart_order.json http://localhost:4000/packageview/orders?url=https://bundles.byte.builders/ui/&name=mongodb-editor-options&version=v0.1.0`

`$ {"kind":"Order","apiVersion":"kubepack.com/v1alpha1","metadata":{"name":"mymongo","namespace":"demo","uid":"d96b7440-2fdb-4fab-89b0-81d2b72631f2","creationTimestamp":"2021-01-13T05:36:37Z"},"spec":{"items":[{"chart":{"url":"https://bundles.byte.builders/ui/","name":"mongodb-editor-options","version":"v0.1.0","releaseName":"mymongo","namespace":"demo","valuesFile":"values.yaml","valuesPatch":[{"op":"add","path":"/metadata/release/name","value":"mymongo"},{"op":"add","path":"/metadata/release/namespace","value":"demo"},{"op":"replace","path":"/spec/version","value":"4.3.2"}]}}]},"status":{}}`

- GET "/packageview/files"

`http://localhost:4000/packageview/files?url=https://bundles.byte.builders/ui/&name=mongodb-editor-options&version=v0.1.0`

- GET "/packageview/files/\*"

`http://localhost:4000/packageview/files/templates/app.yaml?url=https://bundles.byte.builders/ui/&name=mongodb-editor-options&version=v0.1.0`

`http://localhost:4000/packageview/files/values.yaml?url=https://bundles.byte.builders/ui/&name=mongodb-editor-options&version=v0.1.0&format=json`

- PUT "/editor/model" (Initial Model)

`curl -X PUT -H "Content-Type: application/json" -d @./artifacts/mongodb-editor/mongodb_options_model.json http://localhost:4000/editor/model > ./artifacts/mongodb-editor/mongodb_editor_model.json`

- PUT "/editor/manifest" (Preview API)

`curl -X PUT -H "Content-Type: application/json" -d @./artifacts/mongodb-editor/mongodb_editor_model.json http://localhost:4000/editor/manifest > ./artifacts/mongodb-editor/mongodb_editor_manifest.yaml`

- PUT "/editor/resources" (Preview API)

`curl -X PUT -H "Content-Type: application/json" -d @./artifacts/mongodb-editor/mongodb_editor_model.json http://localhost:4000/editor/resources?skipCRDs=true | jq '.' > ./artifacts/mongodb-editor/mongodb_editor_resources.json`

- POST "/deploy/orders"

- GET "/deploy/orders/:id/render/manifest"

http://localhost:4000/deploy/orders/5902b772-319c-40c1-b260-68d81b7864fd/render/manifest

- GET "/deploy/orders/:id/render/resources"
  - Query parameter: skipCRDs=true

http://localhost:4000/deploy/orders/5902b772-319c-40c1-b260-68d81b7864fd/render/resources?skipCRDs=true

- PUT "/clusters/:cluster/editor" (apply/install/update app API)

`curl -X PUT -H "Content-Type: application/json" -d @./artifacts/mongodb-editor/mongodb_editor_model.json  http://localhost:4000/clusters/my_cluster/editor?installCRDs=true`

- DELETE "/clusters/:cluster/editor/namespaces/:namespace/releases/:releaseName" (Delete app api)

`curl -X DELETE -H "Content-Type: application/json" http://localhost:4000/clusters/my_cluster/editor/namespaces/demo/releases/mymongo`

### UI Edit mode

- PUT "/clusters/my_cluster/editor/model"

`curl -X PUT -H "Content-Type: application/json" -d @./artifacts/mongodb-editor/mongodb_editor_model.json  http://localhost:4000/clusters/my_cluster/editor/model`


- GET "/clusters/:cluster/editor/manifest"
  - redundant apis
  - can be replaced by getting the model, then using the /editor apis

`curl -X PUT -H "Content-Type: application/json" -d @./artifacts/mongodb-editor/mongodb_editor_model.json  http://localhost:4000/clusters/my_cluster/editor/manifest`


- GET "/clusters/:cluster/editor/resources"
  - redundant apis
  - can be replaced by getting the model, then using the /editor apis

`curl -X PUT -H "Content-Type: application/json" -d @./artifacts/mongodb-editor/mongodb_editor_model.json  http://localhost:4000/clusters/my_cluster/editor/resources`

## Deploy Button APIs

- http://localhost:4000/products
- http://localhost:4000/products/appscode/kubedb
- http://localhost:4000/product_id/prod_Gnc33bJka9iRl9
- http://localhost:4000/products/appscode/kubedb/plans
- http://localhost:4000/products/appscode/kubedb/compare
- http://localhost:4000/products/appscode/kubedb/plans/kubedb-community
- http://localhost:4000/products/appscode/kubedb/plans/kubedb-community/bundleview

### Generate PackageView

- http://localhost:4000/packageview?url=https://charts.appscode.com/stable/&name=kubedb&version=v0.13.0-rc.0
- http://localhost:4000/packageview?url=https://bundles.kubepack.com&name=stash&version=v0.9.0-rc.6

### Generate BundleView for Chart

- http://localhost:4000/bundleview?url=https://charts.appscode.com/stable/&name=kubedb&version=v0.13.0-rc.0

### Generate order

```console
curl -X POST -H "Content-Type: application/json" -d @artifacts/kubedb-community/bundleview.json http://localhost:4000/bundleview/orders
```

```json
{"kind":"Order","apiVersion":"kubepack.com/v1alpha1","metadata":{"name":"kubedb-community","uid":"1f1d149b-5226-4659-8feb-165face489b3","creationTimestamp":"2020-02-26T12:00:24Z"},"spec":{"items":[{"chart":{"url":"https://charts.appscode.com/stable/","name":"kubedb","version":"v0.13.0-rc.0","releaseName":"kubedb","namespace":"kube-system","bundle":{"name":"kubedb-community","url":"https://bundles.kubepack.com","version":"v0.13.0-rc.0"}}},{"chart":{"url":"https://charts.appscode.com/stable/","name":"kubedb-catalog","version":"v0.13.0-rc.0","releaseName":"kubedb-catalog","namespace":"kube-system","bundle":{"name":"kubedb-community","url":"https://bundles.kubepack.com","version":"v0.13.0-rc.0"}}}]},"status":{}}
```

- http://localhost:4000/deploy/orders/1f1d149b-5226-4659-8feb-165face489b3/helm2
- http://localhost:4000/deploy/orders/1f1d149b-5226-4659-8feb-165face489b3/helm3
- http://localhost:4000/deploy/orders/1f1d149b-5226-4659-8feb-165face489b3/yaml

### List Helm Hub repositories, Charts and Chart Versions

- http://localhost:4000/chartrepositories
- http://localhost:4000/chartrepositories/charts/?url=https://charts.appscode.com/stable/
- http://localhost:4000/chartrepositories/charts/voyager/versions/?url=https://charts.appscode.com/stable/


