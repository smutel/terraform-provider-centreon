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
	go build -o terraform-provider-centreon .

localinstall:
	@echo "==> Creating folder terraform.d/plugins/linux_amd64"
	mkdir -p ~/.terraform.d/plugins/linux_amd64
	@echo "==> Installing provider in this folder"
	cp terraform-provider-centreon ~/.terraform.d/plugins/linux_amd64

check:
	@echo "==> Checking terraform-provider-centreon"
	gometalinter \
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
	rm -f terraform-provider-centreon
