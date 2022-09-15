#!/bin/bash

source "$(dirname "${BASH_SOURCE}")/lib/init.sh"

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${SCRIPT_ROOT}; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../../../k8s.io/code-generator)}

verify="${VERIFY:-}"

${CODEGEN_PKG}/generate-groups.sh "deepcopy" \
  github.com/kube-all/public-cloud/generated \
  github.com/kube-all/public-cloud \
  "aliyun:v1alpha1 tencent:v1alpha1"  \
  --go-header-file ${SCRIPT_ROOT}/hack/boilerplate.txt \
  ${verify}
