PWD:=$(shell pwd)
TARGET=...

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: setup
setup: ## Install depeendent tools and setup project
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.10.1
	go install github.com/volatiletech/sqlboiler/v4@v4.6.0
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.6.0
	go install github.com/cosmtrek/air@latest
	go install github.com/kyoh86/richgo@latest

test: ## Run test
	APP_ENV=test richgo test --cover yawaraka-tissue/${TARGET}

.PHONY: gen-api
gen-api: ## Generate router and request type structs from OpenAPI spec.
	go generate gen/gen_openapi.go

.PHONY: run
run: ## Run API server
	air -c .air.toml
