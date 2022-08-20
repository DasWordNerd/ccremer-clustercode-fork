# Set Shell to bash, otherwise some targets fail with dash/zsh etc.
SHELL := /bin/bash

# Disable built-in rules
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-builtin-variables
.SUFFIXES:
.SECONDARY:

PROJECT_ROOT_DIR = .
include Makefile.vars.mk

e2e_make := $(MAKE) -C e2e
docs_make := $(MAKE) -C docs
media_make := $(MAKE) -C data

go_build ?= CGO_ENABLED=0 go build -o $(BIN_FILENAME) main.go

# Run tests (see https://sdk.operatorframework.io/docs/building-operators/golang/references/envtest-setup)
ENVTEST_ASSETS_DIR=$(shell pwd)/testbin

all: build ## Invokes the build target

.PHONY: test
test: fmt vet ## Run tests
	go test ./...

integration-test: export ENVTEST_K8S_VERSION = $(INTEGRATIONTEST_K8S_VERSION)
integration-test: generate fmt vet $(testbin_created) ## Run integration tests with envtest
	test -f ${ENVTEST_ASSETS_DIR}/setup-envtest.sh || curl -sSLo ${ENVTEST_ASSETS_DIR}/setup-envtest.sh https://raw.githubusercontent.com/kubernetes-sigs/controller-runtime/master/hack/setup-envtest.sh
	source ${ENVTEST_ASSETS_DIR}/setup-envtest.sh; fetch_envtest_tools $(ENVTEST_ASSETS_DIR); setup_envtest_env $(ENVTEST_ASSETS_DIR); go test -tags=integration -v ./... $(test_args)

.PHONY: build
build: generate fmt vet $(BIN_FILENAME) ## Build manager binary

.PHONY: run
run: export CC_OPERATOR__ENABLE_LEADER_ELECTION = $(ENABLE_LEADER_ELECTION)
run: fmt vet ## Run against the configured Kubernetes cluster in ~/.kube/config
	go run ./main.go -v operate

.PHONY: install
install: generate ## Install CRDs into a cluster
	$(KUSTOMIZE) build $(CRD_ROOT_DIR) | kubectl apply $(KIND_KUBECTL_ARGS) -f -

.PHONY: uninstall
uninstall: generate ## Uninstall CRDs from a cluster
	$(KUSTOMIZE) build $(CRD_ROOT_DIR) | kubectl delete -f -

.PHONY: deploy
deploy: generate ## Deploy controller in the configured Kubernetes cluster in ~/.kube/config
	$(KUSTOMIZE) build kustomize/default | kubectl apply -f -

.PHONY: generate
generate: ## Generate manifests e.g. CRD, RBAC etc.
	@go generate -tags=generate ./...

.PHONY: crd
crd: generate ## Generate CRD to file
	$(KUSTOMIZE) build $(CRD_ROOT_DIR) -o $(CRD_FILE)

.PHONY: fmt
fmt: ## Run go fmt against code
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code
	go vet ./...

.PHONY: lint
lint: generate fmt vet ## Invokes the generate, fmt and vet targets
	@echo 'Check for uncommitted changes ...'
	git diff --exit-code

.PHONY: docker-build
docker-build: export GOOS = linux
docker-build: $(BIN_FILENAME) ## Build the docker image
	docker build . -t $(DOCKER_IMG) -t $(QUAY_IMG) -t $(E2E_IMG)

.PHONY: docker-push
docker-push: ## Push the docker image
	docker push $(DOCKER_IMG)
	docker push $(QUAY_IMG)

clean: export KUBECONFIG = $(KIND_KUBECONFIG)
clean: e2e-clean kind-clean media-clean docs-clean ## Cleans up the generated resources
	rm -r testbin/ dist/ bin/ cover.out $(BIN_FILENAME) $(CRD_FILE) || true

.PHONY: help
help: ## Show this help
	@grep -E -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

###
### Assets
###

$(testbin_created):
	mkdir -p $(TESTBIN_DIR)
	# a marker file must be created, because the date of the
	# directory may update when content in it is created/updated,
	# which would cause a rebuild / re-initialization of dependants
	@touch $(testbin_created)

# Build the binary without running generators
.PHONY: $(BIN_FILENAME)
$(BIN_FILENAME):
	$(go_build)

###
### KIND
###

.PHONY: kind-setup
kind-setup: ## Creates a kind instance if one does not exist yet.
	@$(e2e_make) kind-setup

.PHONY: kind-clean
kind-clean: ## Removes the kind instance if it exists.
	@$(e2e_make) kind-clean

.PHONY: kind-e2e-image
kind-e2e-image: docker-build
	$(e2e_make) kind-e2e-image

.PHONY: kind-run
kind-run: export KUBECONFIG = $(KIND_KUBECONFIG)
kind-run: export E2E_TAG = master
kind-run: kind-e2e-image install run ## Runs the operator on the local host but configured for the kind cluster

.PHONY: kind-deploy
kind-deploy: export KUBECONFIG = $(KIND_KUBECONFIG)
kind-deploy: export E2E_TAG = master
kind-deploy: kind-e2e-image deploy

###
### E2E Test
###

.PHONY: e2e-test
e2e-test: export KUBECONFIG = $(KIND_KUBECONFIG)
e2e-test: e2e-setup docker-build install blank-media ## Run the e2e tests
	@$(e2e_make) test

.PHONY: e2e-setup
e2e-setup: export KUBECONFIG = $(KIND_KUBECONFIG)
e2e-setup: ## Run the e2e setup
	@$(e2e_make) setup

.PHONY: e2e-clean-setup
e2e-clean-setup: export KUBECONFIG = $(KIND_KUBECONFIG)
e2e-clean-setup: ## Clean the e2e setup (e.g. to rerun the e2e-setup)
	@$(e2e_make) clean-setup

.PHONY: e2e-clean
e2e-clean: ## Remove all e2e-related resources (incl. all e2e Docker images)
	@$(e2e_make) clean

###
### Documentation
###

docs-clean: ## Remove all documentation resources
	@$(docs_make) clean

docs-preview: ## Preview documentation in local web server and browser
	@$(docs_make) preview

docs-build: ## Build documentation
	@$(docs_make) build

###
### Media
###

media-clean: ## Remove all media resources
	@$(media_make) clean

blank-media: ## Preview documentation in local web server and browser
	@$(media_make) blank-media
