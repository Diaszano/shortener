# Makefile for the Shortener API Project

# Variables
GO_RUN := @go run
MIGRATE := $(GO_RUN) cmd/migrate/migrate.go
DOCKER_COMPOSE := docker compose -f build/docker-compose.yml --env-file .env

# Targets

## Run the main application
run:
	$(GO_RUN) cmd/shortener/shortener.go

## Create a new database migration
## Usage: make migration name=<migration_name>
migration:
	@migrate create -ext sql -dir internal/database/migrations $(filter-out $@, $(MAKECMDGOALS))

## Run database migrations (up)
migrate-up:
	$(MIGRATE) up

## Revert the last database migration (down)
migrate-down:
	$(MIGRATE) down

## Start the Docker environment
docker-up:
	$(DOCKER_COMPOSE) up -d

## Stop the Docker environment
docker-down:
	$(DOCKER_COMPOSE) down

## Generate SQL code using sqlc
sqlcg:
	@sqlc generate

# Prevents errors when calling targets with extra arguments
# Example: make migration name=create_users_table
%:
	@:
