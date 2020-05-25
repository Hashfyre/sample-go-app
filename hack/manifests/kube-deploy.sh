##!/usr/bin/env bash

GIT_TLD=$(git rev-parse --show-toplevel)
set -u

kubectl apply \
  -f ${GIT_TLD}/hack/manifests/namespace.yaml

kubectl apply \
  -f ${GIT_TLD}/hack/manifests/postgres-configmap.yaml \
  -f ${GIT_TLD}/hack/manifests/postgres-sfset.yaml \
  -f ${GIT_TLD}/hack/manifests/postgres-svc.yaml

kubectl apply \
  -f ${GIT_TLD}/hack/manifests/configmap.yaml \
  -f ${GIT_TLD}/hack/manifests/deployment.yaml \
  -f ${GIT_TLD}/hack/manifests/service.yaml
