#!/bin/bash

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

curl 'http://localhost:4000/clusters/appscode/editor?installCRDs=true' \
    -X 'PUT' \
    -H 'Accept: application/json, text/plain, */*' \
    -H 'Content-Type: application/json;charset=UTF-8' \
    --data-raw '{"metadata":{"release":{"name":"db-12","namespace":"demo"},"resource":{"group":"kubedb.com","kind":"Postgres","name":"postgreses","scope":"Namespaced","version":"v1alpha2"}},"resources":{"kubedbComPostgres":{"apiVersion":"kubedb.com/v1alpha2","kind":"Postgres","metadata":{"labels":{"app.kubernetes.io/instance":"db-12","app.kubernetes.io/managed-by":"Helm","app.kubernetes.io/name":"postgreses.kubedb.com"},"name":"db-12","namespace":"demo"},"spec":{"clientAuthMode":"md5","leaderElection":{"electionTick":10,"heartbeatTick":1,"maximumLagBeforeFailover":33554432,"period":"100ms"},"podTemplate":{"spec":{"resources":{"limits":{"cpu":"1","memory":"1024Mi"}}}},"replicas":3,"standbyMode":"Hot","storage":{"accessModes":["ReadWriteOnce"],"resources":{"requests":{"storage":"10Gi"}},"storageClassName":"linode-block-storage"},"storageType":"Durable","terminationPolicy":"WipeOut","version":"9.6.21-debian","sslMode":"require","tls":{"issuerRef":{"apiGroup":"cert-manager.io","kind":"Issuer","name":"postgres-ca-issuer"}},"monitor":{"agent":"prometheus.io/operator","prometheus":{"serviceMonitor":{"labels":{"release":"prometheus"}}}},"configSecret":{"name":"db-12-config"}}},"secret_config":{"stringData":{"user.conf":"max_connections = 121\nshared_buffers = 121MB"}}},"spec":{"version":"9.6.21-debian"}}'

curl 'http://localhost:4000/clusters/appscode/editor?installCRDs=true' \
    -X 'PUT' \
    -H 'Accept: application/json, text/plain, */*' \
    -H 'Content-Type: application/json;charset=UTF-8' \
    --data-raw '{"metadata":{"release":{"name":"db-12","namespace":"demo"},"resource":{"group":"kubedb.com","kind":"Postgres","name":"postgreses","scope":"Namespaced","version":"v1alpha2"}},"patch":[{"op":"remove","path":"/resources/kubedbComPostgres/spec/podTemplate/spec/containerSecurityContext/capabilities"},{"op":"remove","path":"/resources/kubedbComPostgres/spec/podTemplate/metadata"},{"op":"remove","path":"/resources/kubedbComPostgres/spec/podTemplate/controller"},{"op":"remove","path":"/resources/kubedbComPostgres/spec/monitor/prometheus/exporter/resources"}]}'
