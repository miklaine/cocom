help:
	@echo "Please use \`make <ROOT>' where <ROOT> is one of"
	@echo "  update-dependencies  to update dependencies"
	@echo "  dependencies         to install the dependencies"
	@echo "  build             	  to build the main binary for current platform"
	@echo "  clean                to remove generated files"


update-dependencies:
	go get -u ./...

dependencies:
	go mod download

clean:
	rm -rf ./build

build: $(GO_PACKAGES)
	$(GO_VARS) $(GO) build -o="cocom" -ldflags="$(LD_FLAGS)" $(ROOT)./cmd

.PHONY: update-dependencies dependencies build clean

GO_VARS ?=
GO_PACKAGES := $(shell go list ./...)
GO ?= go
GIT ?= git
LD_FLAGS := -s -w
TOPDIR=$(PWD)
