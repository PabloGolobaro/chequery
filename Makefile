export DB_PORT := 5432
export DB_HOST := localhost
export DB_USER := golobar
export DB_PASS := password
export DB_NAME := golo
export SERVER_HOST := localhost
export SERVER_PORT := 80
export JWT_SECRET := secret

#VERSION := $(shell git describe --tags --always --dirty)
#
#REGISTRY := pablogolobar/order_server:$(VERSION)

migrate-status:
	goose postgres " host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable" status

migrate: migrate-status
	goose -dir ./migrations postgres " host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable" up

up:
	docker-compose up -d
server: up migrate
	go run ./cmd/server
client:
	go run ./cmd/client

spec:
	swagger generate spec -m -w ./cmd/server -o ./api/swagger.yaml
	swagger generate spec -m -w ./cmd/server -o ./static/swagger/swagger.json

swagger: spec
	swagger serve -F=swagger ./api/swagger.yaml

lint:
	golangci-lint run ./internal/... -v
test:
	go test ./internal/...
	go test ./pkg/...
build:  test
	go build -o ./bin/server.exe ./cmd/server
	go build -o ./bin/client.exe ./cmd/client


path:
	path D:\Go\Swagger;%PATH%
client-gen: spec
	swagger generate client -f ./api/swagger.yaml -t ./cmd/client
