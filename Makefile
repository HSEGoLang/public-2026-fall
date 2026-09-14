GO_TOOLCHAIN_VERSION := go1.27.1
GOLANGCI_LINT_VERSION := v2.13.2
GOLANGCI_LINT := $(shell go env GOPATH)/bin/golangci-lint
GO_MODULES := . lectures lectures/12-lowlevel/unix 15-Bonus/wasm/flappygopher

export GOTOOLCHAIN := $(GO_TOOLCHAIN_VERSION)

-include Makefile.private

HELP_FUN = \
	%help; while(<>){push@{$$help{$$2//'options'}},[$$1,$$3] \
	if/^([\w-_]+)\s*:.*\#\#(?:@(\w+))?\s(.*)$$/}; \
    print"$$_:\n", map"  $$_->[0]".(" "x(20-length($$_->[0])))."$$_->[1]\n",\
    @{$$help{$$_}},"\n" for keys %help; \

args := $(wordlist 2, 100, $(MAKECMDGOALS))
ifndef args
MESSAGE = "No such command (or you pass two or many targets to ). List of possible commands: make help"
else
MESSAGE = "Done"
endif

# Commands

.PHONY: help
help: ##@Help Show this help
	@echo -e "Usage: make [target] ...\n"
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)

.PHONY: setup
setup: download-deps $(GOLANGCI_LINT) ##@Environment Download dependencies and install development tools
	@go version
	@$(GOLANGCI_LINT) --version

.PHONY: download-deps
download-deps: ##@Environment Download dependencies for all Go modules
	@set -e; for module in $(GO_MODULES); do \
		if [ ! -f "$$module/go.mod" ]; then \
			echo "Skipping $$module: go.mod not found"; \
			continue; \
		fi; \
		echo "Downloading Go dependencies in $$module"; \
		go -C "$$module" mod download; \
	done

$(GOLANGCI_LINT):
	curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b "$$(go env GOPATH)/bin" "$(GOLANGCI_LINT_VERSION)"

.PHONY: format
format: | $(GOLANGCI_LINT) ###@Code Format files
	$(GOLANGCI_LINT) run --no-config --enable-only wsl_v5 --fix ./$(args)/...
	$(GOLANGCI_LINT) fmt ./$(args)/...

.PHONY: rebase-on-public
rebase-on-public: ##@Git Rebase on public repo
	git remote add public-repo git@gitlab.manytask.org:hse/go/public-2026-fall.git || true
	git fetch public-repo
	git rebase public-repo/main

%::
	echo $(MESSAGE)
