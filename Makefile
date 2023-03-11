export DB_PORT := 5432
export DB_HOST := localhost
export DB_USER := golobar
export DB_PASS := password
export DB_NAME := golo
export SERVER_HOST := localhost
export SERVER_PORT := 8080

#VERSION := $(shell git describe --tags --always --dirty)
#
#REGISTRY := pablogolobar/order_server:$(VERSION)

up:
	docker-compose up -d

run: up
	go run ./cmd/server
spec:
	swagger generate spec -m -w ./cmd/server -o ./api/swagger.yaml
serve:
	swagger serve -F=swagger ./api/swagger.yaml
lint:
	@GO111MODULE=on golangci-lint run ./metr-checker/... -v
