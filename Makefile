.PHONY: help build build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## Build docker image to deploy
	docker build -t tatuya-web/go-gin-template:${DOCKER_TAG} \
		--target deploy ./

build-local: ## Build docker image to local development
	docker compose build --no-cache

up: ## Do docker compose up with hot reload
	docker compose up -d

down: ## Do docker compose down
	docker compose down

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check container status
	docker compose ps

test: ## Execute tests
	go test -race -shuffle=on ./...

dry-migrate: ## Try migration
	mysqldef -u db_user -p db_password -h 127.0.0.1 -p 33306 go_gin_template --dry-run < ./_tools/mysql/schema.sql

migrate:  ## Execute migration
	mysqldef -u db_user -p db_password -h 127.0.0.1 -p 33306 go_gin_template < ./_tools/mysql/schema.sql

generate: ## Generate codes
	go generate ./...

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# mysqldef -u db_user -p db_password -h db go_gin_template < ./_tools/mysql/schema.sql
# mysqldef -u db_user -p db_password -h db go_gin_template --dry-run < ./_tools/mysql/schema.sql
# go install github.com/k0kubun/sqldef/cmd/mysqldef@latest