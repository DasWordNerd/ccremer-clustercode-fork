IMG_TAG ?= latest

BIN_FILENAME ?= $(PROJECT_ROOT_DIR)/clustercode
TESTBIN_DIR ?= $(PROJECT_ROOT_DIR)/testbin/bin

# See https://storage.googleapis.com/kubebuilder-tools/ for list of supported K8s versions
INTEGRATIONTEST_K8S_VERSION ?= 1.20.2

KIND_VERSION ?= 0.9.0
KIND_NODE_VERSION ?= v1.20.0
KIND ?= $(TESTBIN_DIR)/kind

KIND_KUBECONFIG ?= $(TESTBIN_DIR)/kind-kubeconfig-$(KIND_NODE_VERSION)
KIND_CLUSTER ?= clustercode-$(KIND_NODE_VERSION)
KIND_KUBECTL_ARGS ?= --validate=true

E2E_TAG ?= e2e_$(shell sha1sum $(BIN_FILENAME) | cut -b-8)
E2E_REPO ?= local.dev/clustercode/e2e
E2E_IMG = $(E2E_REPO):$(E2E_TAG)
OPERATOR_NAMESPACE ?= clustercode-system

# Image URL to use all building/pushing image targets
CONTAINER_IMG ?= ghcr.io/ccremer/clustercode:$(IMG_TAG)
FFMPEG_IMG ?= docker.io/jrottenberg/ffmpeg:4.1-alpine

testbin_created = $(TESTBIN_DIR)/.created
