# Load environment variables from .env file
include .env
export

.PHONY: all build run migrate clean

all: build run

run-dev:
	go run ./main.go

build:
	go build -o arvigo .

clean:
	rm -f arvigo

migrate-up:
	migrate -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -source=file://./pkg/database/migration up

migrate-down:
	migrate -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -source=file://./pkg/database/migration down
