SHELL := /usr/bin/env bash
.DEFAULT_GOAL := help
KUBECONFIG := ${HOME}/.kube/config

## --------------------------------------
## Help
## --------------------------------------
##@ help:

help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
## Tooling
## --------------------------------------
##@ tooling:

.PHONY: install-carvel
install-carvel:  ## Install carvel tools on bin folder
	 curl -L https://carvel.dev/install.sh | K14SIO_INSTALL_BIN_DIR=bin $(SHELL)

.PHONY: install-kind
install-kind: install-carvel ## Install kind and Kapp CRDs
	kind delete cluster || true; \
	./scripts/kind-with-registry.sh; \
	bin/kapp deploy -a kc -f https://github.com/carvel-dev/kapp-controller/releases/latest/download/release.yml

## --------------------------------------
## Running
## --------------------------------------
##@ running:

.PHONY: run-local
run-local:  ## Run server as a local binary
	go run . --kubeconfig=$(KUBECONFIG)

.PHONY: react-dev
react-dev:  ## Run local frontend server
	cd attestation; npm run start

.PHONY: react-build
react-build:  ## Run local frontend server
	cd attestation; npm run build
