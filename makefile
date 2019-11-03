SHELL := /bin/bash
MAKEFLAGS += --warn-undefined-variables
.DEFAULT_GOAL := build
.PHONY: *

tools: ## Download and install all dev/code tools
	@echo "==> Installing dev tools"
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

build:
	@echo "==> Building terraform-provider-centreon"
	@go build -o terraform-provider-centreon .

build-tarball: clean
	@echo "==> Building terraform-provider-centreon"
	@mkdir -p builds
	@if [[ -n $${version} ]]; then export VERSION="_$${version}"; fi ; for GOOS in darwin linux windows; do for GOARCH in 386 amd64; do env GOOS=$${GOOS} GOARCH=$${GOARCH} go build -o builds/terraform-provider-centreon-$${GOOS}-$${GOARCH}$${VERSION} .; done; done
	@cd builds && for file in *; do echo "$$(sha256sum $${file})" >> ../terraform-provider-centreon_SHA256SUM; done
	@tar czf terraform-provider-centreon.tar.gz builds

localinstall:
	@echo "==> Creating folder terraform.d/plugins/linux_amd64"
	@mkdir -p ~/.terraform.d/plugins/linux_amd64
	@echo "==> Installing provider in this folder"
	@cp terraform-provider-centreon ~/.terraform.d/plugins/linux_amd64

check:
	@echo "==> Checking terraform-provider-centreon"
	@gometalinter \
		--deadline 10m \
		--vendor \
		--sort="path" \
		--aggregate \
		--enable-gc \
		--disable-all \
		--enable goimports \
		--enable misspell \
		--enable vet \
		--enable deadcode \
		--enable varcheck \
		--enable ineffassign \
		--enable gofmt \
		--enable golint \
		./...

clean:
	@echo "==> Cleaning terraform-provider-centreon"
	@rm -f terraform-provider-centreon
	@rm -rf builds
	@rm -f terraform-provider-centreon_SHA256SUM
	@rm -f terraform-provider-centreon.tar.gz
