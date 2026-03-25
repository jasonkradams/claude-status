default: help

##@ Build

WIDGETS    := multi-line
BINS       := $(addprefix bin/,$(WIDGETS))
GO_SOURCES := $(shell find . -name '*.go' -not -path './vendor/*')

build: $(BINS) ## build all widgets to bin/

.prepare: $(GO_SOURCES)
	@echo "==> Formatting..."
	@go fmt ./...
	@echo "==> Vetting..."
	@go vet ./...
	@echo "==> Fixing..."
	@go fix ./...
	@touch $@

$(BINS): bin/%: .prepare
	@"$(CURDIR)/scripts/build-widget.sh" $* ./cmd/$*

build-%: bin/%

##@ Utilities

install: build ## install a widget as ~/bin/claude-status
	@"$(CURDIR)/scripts/install.sh"

clean: ## clean build artifacts
	@"$(CURDIR)/scripts/clean.sh"

help:
	@awk ' \
		BEGIN { \
			green = "\033[32m"; \
			bold  = "\033[1m"; \
			reset = "\033[0m"; \
		} \
		/^##@/ { \
			printf "\n%s%s%s\n", bold, substr($$0, 5), reset; \
			next; \
		} \
		/^[a-zA-Z_-]+:.*?## / { \
			split($$0, a, ":.*?## "); \
			printf "  %s%-14s%s %s\n", green, a[1], reset, a[2]; \
		} \
	' $(MAKEFILE_LIST)

.NOTPARALLEL:

.PHONY: build $(addprefix build-,$(WIDGETS)) install clean help default
