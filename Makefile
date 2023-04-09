SHELL := /usr/bin/env bash

.DEFAULT_GOAL := help

# ##############################################################################
# Makefile targets/goals
# ##############################################################################

##
## Help commands
## -----
##

.PHONY: list
list: ## Generate basic list of all targets
	@grep '^[^\.#[:space:]].*:' Makefile | \
		grep -v -e "[:?]=" -e "%$\" | \
		cut -d':' -f1

.PHONY: help
help: ## Makefile help
	@grep -E '(^[a-zA-Z_0-9%/-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | \
		sed -e 's/\[32m##/[33m/'

##
## Main commands
## -----
##

.PHONY: synchronize
synchronize: ## Synchronize module's dependencies
	go mod tidy

.PHONY: run
run: guard-SLACK_TOKEN guard-SLACK_URL ## Launch the Slack Emoji Downloader
	go run .

.PHONY: test
test: ## Launch the tests
	go test -v

.PHONY: build
build: ## Build project
	go build -v ./...

##
## Misc commands
## -----
##

.PHONY: guard-%
guard-%: ## Check if a given env var is well set, usage: make guard-SLACK_TOKEN
	@if [ "${${*}}" = "" ]; then \
		echo -e "$(COLOR_RED)Problem :: $* should be exported in your env$(COLOR_RESET)"; \
		exit 1; \
	fi
