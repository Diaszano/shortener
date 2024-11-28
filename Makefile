run:
	@go run cmd/shortener/main.go

migration:
	@migrate create -ext sql -dir db/migrations $(filter-out $@, $(MAKECMDGOALS))

DOCKER_COMPOSE := docker compose -f build/docker-compose.yml --env-file .env

docker-up:
	@$(DOCKER_COMPOSE) up -d

docker-down:
	@$(DOCKER_COMPOSE) down

sqlc-generate:
	sqlc generate
