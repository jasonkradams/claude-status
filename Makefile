default: help

##@ Build

GO_SOURCES := $(shell find . -name '*.go' -not -path './vendor/*')

build: bin/multi-line ## build all widgets to bin/

build-multi-line: bin/multi-line

bin/multi-line: $(GO_SOURCES)
	@"$(CURDIR)/scripts/build-multi-line.sh"

##@ Utilities

install: build ## install a widget as ~/bin/claude-status
	@"$(CURDIR)/scripts/install.sh"

clean: ## clean build artifacts
	@"$(CURDIR)/scripts/clean.sh"

help:
	@awk ' \
		BEGIN { \
			cyan  = "\033[36m"; \
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

.PHONY: build build-multi-line install clean help default
