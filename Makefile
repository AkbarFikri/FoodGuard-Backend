include .env

build:
	@echo 'Build app...'
	@go build -o bin/foodguard cmd/app/main.go

run: build
	@echo 'Starting app...'
	@./bin/foodguard

migrate-create:
	@echo 'Creating Migrations...'
	@migrate create -ext sql -dir database/migrations $(name)

migrate-up:
	@migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" -path database/migrations up

migrate-down:
	@migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" -path database/migrations down