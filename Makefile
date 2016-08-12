TRAVIS_BRANCH   = $(shell echo $$TRAVIS_BRANCH)
TRAVIS_TAG      = $(shell echo $$TRAVIS_TAG)
GITHUB_TOKEN    = $(shell echo $$GITHUB_TOKEN)
GOX_BINARY_PATH = dist/{{.Dir}}_{{.OS}}_{{.Arch}}
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
	@echo No Travis CI >&2
	@exit 1
endif

travis-test: assert-on-travis gom install-test-dependencies test ## Run tests in Travis CI

travis-release: assert-on-travis release ## Release binaries on GitHub by Travis CI

release: ## Release binaries on GitHub by the specified tag
ifeq ($(TRAVIS_TAG),)
	@echo Skip the release process because TRAVIS_TAG is empty
else
	go get github.com/mitchellh/gox
	@echo cross compile
	gox -output '$(GOX_BINARY_PATH)' -os '$(GOX_OS)' -arch '$(GOX_ARCH)'
	@echo archive each binary
	for i in dist/*; do j=$$(echo $$i | sed -e 's/_[^\.]*//'); mv $$i $$j; zip -j $${i%.*}.zip $$j; rm $$j; done
	go get github.com/tcnksm/ghr
	@echo Releasing binaries on tag: $(TRAVIS_TAG)
	@ghr --username delphinus35 --token $(GITHUB_TOKEN) --replace --prerelease --debug $(TRAVIS_TAG) dist/
endif
