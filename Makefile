.PHONY: default help

default: help
help: ## help: display make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m make %-20s -> %s\n\033[0m", $$1, $$2}'

# make: app info
APP_NAME     := magic
APP_WORKDIR  := $(shell pwd)
APP_PACKAGES := $(shell go list -f '{{.Dir}}' ./...)
APP_LOG_FMT  := `/bin/date "+%Y-%m-%d %H:%M:%S %z [$(APP_NAME)]"`

# --------------------------------------------------
# Runtime Targets
# --------------------------------------------------
.PHONY: up
up: ## runtime: start local environment
	@echo $(APP_LOG_FMT) "starting local environment"
	@docker compose up --build --remove-orphans --detach

.PHONY: status
status: ## runtime: check local environment status
	@echo $(APP_LOG_FMT) "checking environment status"
	@docker compose ps \
		&& docker compose logs api

.PHONY: down
down: ## runtime: stop local environment
	@echo $(APP_LOG_FMT) "stopping local environment"
	@docker compose down -v --rmi local --remove-orphans

.PHONY: restart
restart: down up ## runtime: restart environment

# --------------------------------------------------
# Build Targets
# --------------------------------------------------
BUILD_DIR := $(APP_WORKDIR)/build

.PHONY: build-clean
build-clean: ## build: clean build workspace
	@echo $(APP_LOG_FMT) "cleaning build workspace"
	@rm -rf $(BUILD_DIR)

.PHONY: build-binary
build-binary: build-clean ## build: build binary file
	@echo $(APP_LOG_FMT) "building binary"
	@mkdir -p $(BUILD_DIR)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build \
		-o $(BUILD_DIR)/cardmodd -ldflags '-extldflags "-static"' \
		cmd/cardmodd/main.go

.PHONY: build-proto
build-proto: ## build: generate proto files and swagger docs
	@echo $(APP_LOG_FMT) "generating proto files and swagger docs"
	@buf generate $(APP_WORKDIR)/internal/proto

# --------------------------------------------------
# Test Targets
# --------------------------------------------------
COVERAGE_DIR := $(APP_WORKDIR)/coverage
LINT_DIR     := $(COVERAGE_DIR)/lint
UNIT_DIR     := $(COVERAGE_DIR)/unit

# unit coverage
UNIT_WEBPAGE  := $(UNIT_DIR)/index.html
UNIT_REPORT   := $(UNIT_DIR)/report.out
UNIT_COVERAGE := $(UNIT_DIR)/coverage.out

.PHONY: test-clean
test-clean: ## test: clean test workspace
	@echo $(APP_LOG_FMT) "cleaning workspace"
	@rm -rf $(COVERAGE_DIR)

.PHONY: test-lint
test-lint: ## test: check for lint failures
	@echo $(APP_LOG_FMT) "checking for lint failures"
	@golangci-lint run -v

.PHONY: test-unit
test-unit: ## test: execute unit test suite
	@echo $(APP_LOG_FMT) "executing unit test suite"
	@mkdir -p $(UNIT_DIR)
	@go test -v \
		-covermode=atomic \
		-coverprofile=$(UNIT_COVERAGE) \
		$(APP_PACKAGES) \
		2>&1 > $(UNIT_REPORT) || cat $(UNIT_REPORT)
	@cat $(UNIT_REPORT)
	@go tool cover -func=$(UNIT_COVERAGE)
	@go tool cover -html=$(UNIT_COVERAGE) -o $(UNIT_WEBPAGE)