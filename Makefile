

verify:
	chmod a+x  ./vendor/k8s.io/code-generator/generate-groups.sh
	hack/verify-deepcopy.sh
	hack/verify-protobuf.sh
	hack/verify-codegen.sh
	hack/verify-swagger-docs.sh

.PHONY: verify


.PHONY: generate
generate:
	chmod a+x  ./vendor/k8s.io/code-generator/generate-groups.sh
	hack/update-deepcopy.sh
	hack/update-protobuf.sh
	hack/update-swagger-docs.sh
	hack/update-codegen.sh




CONTROLLER_GEN = $(GOPATH)/bin/controller-gen

.PHONY: manifests
manifests:
	controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=manifests/crd

