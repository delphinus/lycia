DIST_DIR        = dist
GOX_BINARY_PATH = $(DIST_DIR)/{{.Dir}}_{{.OS}}_{{.Arch}}
GOX_OS          = darwin linux windows
GOX_ARCH        = 386 amd64

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
	$(error No Travis CI)
endif

assert-with-travis-tag: ## Assert with $TRAVIS_TAG
ifeq ($(TRAVIS_TAG),)
	$(error No TRAVIS_TAG environmental variable)
endif

travis-test: assert-on-travis gom install-test-dependencies test ## Run tests in Travis CI

travis-release: assert-on-travis release ## Release binaries on GitHub by Travis CI

release: ## Release binaries on GitHub by the specified tag
ifeq ($(TRAVIS_TAG),)
	$(warning No TRAVIS_TAG environmental variable)
else
	$(call cross-compile)
	: Releasing binaries on tag: $(TRAVIS_TAG)
	go get github.com/tcnksm/ghr
	@ghr --username delphinus35 --token $(GITHUB_TOKEN) --replace --prerelease --debug $(TRAVIS_TAG) dist/
endif

build-all: ## Build all binaries in /dist
	$(call cross-compile)

# $(call cross-compile)
define cross-compile
	rm -fr $(DIST_DIR)
	: cross compile
	go get github.com/mitchellh/gox
	gox -output '$(GOX_BINARY_PATH)' -os '$(GOX_OS)' -arch '$(GOX_ARCH)'
	: archive each binary
	for i in dist/*; \
	do \
		j=$$(echo $$i | sed -e 's/_[^\.]*//'); \
		mv $$i $$j; \
		zip -j $${i%.*}.zip $$j; \
		rm $$j; \
	done
endef
