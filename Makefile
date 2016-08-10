.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test-all: gom package-install-test test ## Run tests with installing `gom` & other dependencies

test: ## Run tests only
	gom test `go list ./... | grep -v vendor`

build: gom package-install ## Build binary
	go build

gom: ## Install gom
	go get github.com/mattn/gom

gom-update: ## Update gom
	go get -u github.com/mattn/gom

package-install: ## Install packages
	gom install

package-install-test: ## Install packages with the `test` group
	gom -test install
