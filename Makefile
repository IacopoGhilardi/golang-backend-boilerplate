.PHONY: build test run migrate docker-up docker-down

# Variabili
APP_NAME=myapp
DOCKER_COMPOSE=docker-compose

build:
	go build -o $(APP_NAME) ./cmd/$(APP_NAME)

test:
	go test -v ./...

run:
	go run ./cmd/$(APP_NAME)

migrate:
	go run ./cmd/migrate

docker-up:
	$(DOCKER_COMPOSE) up -d

docker-down:
	$(DOCKER_COMPOSE) down

lint:
	golangci-lint run

generate-swagger:
	swag init -g cmd/$(APP_NAME)/main.go
