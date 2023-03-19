export DB_PORT := 5432
export DB_HOST := localhost
export DB_USER := golobar
export DB_PASS := password
export DB_NAME := golo
export SERVER_HOST := localhost
export SERVER_PORT := 80

#VERSION := $(shell git describe --tags --always --dirty)
#
#REGISTRY := pablogolobar/order_server:$(VERSION)

up:
	docker-compose up -d
dev: up
	go run ./cmd/server

spec:
	swagger generate spec -m -w ./cmd/server -o ./api/swagger.yaml
swagger: spec
	swagger serve -F=swagger ./api/swagger.yaml

lint:
	@GO111MODULE=on golangci-lint run ./metr-checker/... -v
test:
	go test ./...
build: lint test
	go build ./cmd/server

path:
	path D:\Go\Swagger;%PATH%
client-gen: spec
	swagger generate client -f ./api/swagger.yaml -t ./cmd/client


migrate-status:
	goose postgres " host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable" status

migrate: migrate-status
	goose -dir ./migrations postgres " host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable" up
