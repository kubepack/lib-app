# Copyright AppsCode Inc. and Contributors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

curl 'http://localhost:4000/clusters/console-demo-linode/editor?installCRDs=true' \
  -X 'PUT' \
  -H 'Accept: application/json, text/plain, /' \
  -H 'Content-Type: application/json' \
  --data-raw '{"form":{"alert":{"additionalRuleLabels":{},"annotations":{},"enabled":true,"groups":{"cluster":{"enabled":true,"rules":{"galeraReplicationLatencyTooLong":{"duration":"5m","enabled":true,"severity":"warning","val":0.1}}},"database":{"enabled":true,"rules":{"mysqlHighIncomingBytes":{"duration":"0m","enabled":true,"severity":"critical","val":1048576},"mysqlHighOutgoingBytes":{"duration":"0m","enabled":true,"severity":"critical","val":1048576},"mysqlHighQPS":{"duration":"0m","enabled":true,"severity":"critical","val":1000},"mysqlHighThreadsRunning":{"duration":"2m","enabled":true,"severity":"warning","val":60},"mysqlInnoDBLogWaits":{"duration":"0m","enabled":true,"severity":"warning","val":10},"mysqlInstanceDown":{"duration":"0m","enabled":true,"severity":"critical"},"mysqlRestarted":{"duration":"0m","enabled":true,"severity":"warning","val":60},"mysqlServiceDown":{"duration":"0m","enabled":true,"severity":"critical"},"mysqlSlowQueries":{"duration":"2m","enabled":true,"severity":"warning"},"mysqlTooManyConnections":{"duration":"2m","enabled":true,"severity":"warning","val":80},"mysqlTooManyOpenFiles":{"duration":"2m","enabled":true,"severity":"warning","val":80}}},"opsManager":{"enabled":true,"rules":{"opsRequestFailed":{"duration":"0m","enabled":true,"severity":"critical"},"opsRequestOnProgress":{"duration":"0m","enabled":true,"severity":"info"},"opsRequestStatusProgressingToLong":{"duration":"30m","enabled":true,"severity":"critical"}}},"provisioner":{"enabled":true,"rules":{"appPhaseCritical":{"duration":"15m","enabled":true,"severity":"warning"},"appPhaseNotReady":{"duration":"1m","enabled":true,"severity":"critical"}}},"schemaManager":{"enabled":true,"rules":{"schemaExpired":{"duration":"0m","enabled":true,"severity":"warning"},"schemaFailed":{"duration":"0m","enabled":true,"severity":"warning"},"schemaInProgressForTooLong":{"duration":"30m","enabled":true,"severity":"warning"},"schemaPendingForTooLong":{"duration":"30m","enabled":true,"severity":"warning"},"schemaTerminatingForTooLong":{"duration":"30m","enabled":true,"severity":"warning"}}},"stash":{"enabled":true,"rules":{"backupSessionFailed":{"duration":"0m","enabled":true,"severity":"critical"},"backupSessionPeriodTooLong":{"duration":"0m","enabled":true,"severity":"warning","val":1800},"noBackupSessionForTooLong":{"duration":"0m","enabled":true,"severity":"warning","val":18000},"repositoryCorrupted":{"duration":"5m","enabled":true,"severity":"critical"},"repositoryStorageRunningLow":{"duration":"5m","enabled":true,"severity":"warning","val":10737418240},"restoreSessionFailed":{"duration":"0m","enabled":true,"severity":"critical"},"restoreSessionPeriodTooLong":{"duration":"0m","enabled":true,"severity":"warning","val":1800}}}},"labels":{"release":"kube-prometheus-stack"}}},"metadata":{"release":{"name":"mariadb-test","namespace":"demo"},"resource":{"group":"kubedb.com","kind":"MariaDB","name":"mariadbs","scope":"Namespaced","version":"v1alpha2"}},"resources":{"kubedbComMariaDB":{"apiVersion":"kubedb.com/v1alpha2","kind":"MariaDB","metadata":{"labels":{"app.kubernetes.io/instance":"mariadb-test","app.kubernetes.io/managed-by":"Helm","app.kubernetes.io/name":"mariadbs.kubedb.com"},"name":"mariadb-test","namespace":"demo"},"spec":{"monitor":{"agent":"prometheus.io/operator","prometheus":{"exporter":{"resources":{"requests":{"cpu":"100m","memory":"128Mi"}}},"serviceMonitor":{"interval":"30s","labels":{}}}},"podTemplate":{"spec":{"resources":{"limits":{"cpu":".5","memory":"1024Mi"}}}},"replicas":1,"storage":{"accessModes":["ReadWriteOnce"],"resources":{"requests":{"storage":"10Gi"}},"storageClassName":"linode-block-storage-retain"},"storageType":"Durable","terminationPolicy":"WipeOut","version":"10.5.8"}}}}'

curl 'http://localhost:4000/clusters/console-demo-linode/editor/model' \
  -X 'PUT' \
  -H 'Accept: application/json, text/plain, /' \
  -H 'Content-Type: application/json' \
  --data-raw '{"metadata":{"release":{"name":"mariadb-test","namespace":"demo"},"resource":{"group":"kubedb.com","version":"v1alpha2","name":"mariadbs","resourceTitle":"MariaDB","scope":"Namespaced"}}}'


> cd ../kube-ui-server (master)
> k apply -f charts/kube-ui-server/crds
> k create -f ui.yaml

apiVersion: source.toolkit.fluxcd.io/v1beta2
kind: HelmRepository
metadata:
  name: bytebuilders-ui
  namespace: kubeops
spec:
  interval: 30m
  url: https://bundles.byte.builders/ui/

> k apply -f ~/go/src/kubedb.dev/apimachinery/crds

> k apply -f ~/go/src/github.com/appscode/alerts/charts/mariadb-alerts/crds
