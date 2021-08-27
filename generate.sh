#!/bin/bash

go run cmd/gen-simple-editor/main.go --all --skipExisting=false

go run cmd/fuse-chart/*.go \
  --sample-dir=/home/tamal/go/src/github.com/appscode/kubedb-samples/mongodb/sharded/prometheus.io/backupconfiguration/stash/tls/custom-auth/sharded \
  --chart-dir=/home/tamal/go/src/go.bytebuilders.dev/ui-wizards/charts \
  --sample-name=mongodb \
  --resource.group=kubedb.com \
  --resource.version=v1alpha2 \
  --resource.name=mongodbs

go run cmd/fuse-chart/*.go \
  --sample-dir=/home/tamal/go/src/github.com/appscode/kubedb-samples/postgres/cluster/prometheus.io/backupconfiguration/stash/tls/custom-auth/custom-auth-mode/custom-config/custom-pg-coordinator/custom-uid \
  --chart-dir=/home/tamal/go/src/go.bytebuilders.dev/ui-wizards/charts \
  --sample-name=postgres \
  --resource.group=kubedb.com \
  --resource.version=v1alpha2 \
  --resource.name=postgreses

go run cmd/fuse-chart/*.go \
  --sample-dir=/home/tamal/go/src/github.com/appscode/kubedb-samples/mariadb/cluster/prometheus.io/backupconfiguration/stash/tls/custom-auth/config-file/customize-pod-template \
  --chart-dir=/home/tamal/go/src/go.bytebuilders.dev/ui-wizards/charts \
  --sample-name=mariadb \
  --resource.group=kubedb.com \
  --resource.version=v1alpha2 \
  --resource.name=mariadbs

go run cmd/fuse-chart/*.go \
  --sample-dir=/home/tamal/go/src/github.com/appscode/kubedb-samples/elasticsearch/topology/prometheus.io/backupconfiguration/stash/cert-manager/custom-auth/custom-config/internal-users/kernel-settings \
  --chart-dir=/home/tamal/go/src/go.bytebuilders.dev/ui-wizards/charts \
  --sample-name=elasticsearch \
  --resource.group=kubedb.com \
  --resource.version=v1alpha2 \
  --resource.name=elasticsearches

go run cmd/fuse-chart/*.go \
  --sample-dir=/home/tamal/go/src/github.com/appscode/kubedb-samples/mysql/group-replication/prometheus.io/backupconfiguration/stash/tls/custom-auth/config-file/customize-pod-template \
  --chart-dir=/home/tamal/go/src/go.bytebuilders.dev/ui-wizards/charts \
  --sample-name=mysql \
  --resource.group=kubedb.com \
  --resource.version=v1alpha2 \
  --resource.name=mysqls
