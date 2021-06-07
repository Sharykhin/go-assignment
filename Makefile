.PHONY: dev test build

SERVER_PORT ?= 3000

dev:
	SERVER_PORT=$(SERVER_PORT) go run -race cmd/server/main.go

test:
	go test -v -count=1 ./switcher

build:
	docker build -t go-assignment .

start: build
	docker run --rm -p $(SERVER_PORT):3000 go-assignment
