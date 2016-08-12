TRAVIS_BRANCH   = $(shell echo $$TRAVIS_BRANCH)

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

gom: ## Install gom
	go get github.com/mattn/gom

gom-update: ## Update gom
	go get -u github.com/mattn/gom

test: ## Run tests only
	gom test `go list ./... | grep -v vendor`

build: gom install-dependencies ## Build binary
	go build

install-dependencies: ## Install packages for dependencies
	gom install

install-test-dependencies: ## Install packages for dependencies with the `test` group
	gom -test install

assert-on-travis: ## Assert that this task is executed in Travis CI
ifeq ($(TRAVIS_BRANCH),)
	@echo No Travis CI >&2
	@exit 1
endif

travis-test: assert-on-travis gom install-test-dependencies test ## Run tests in Travis CI
