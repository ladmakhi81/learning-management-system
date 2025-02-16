export DB_USERNAME=$(shell cat .env | grep DB_USER | cut -d '=' -f2)
export DB_PASSWORD=$(shell cat .env | grep DB_PASSWORD | cut -d '=' -f2)
export DB_HOST=$(shell cat .env | grep DB_HOST | cut -d '=' -f2)
export DB_PORT=$(shell cat .env | grep DB_PORT | cut -d '=' -f2)
export DB_NAME=$(shell cat .env | grep DB_NAME | cut -d '=' -f2)


build-server:
	@go build -o ./bin/app ./cmd/api/main.go
run-server:build-server
	@./bin/app
create-migration:
	@migrate create -ext sql -dir migrations $(name)
run-migration:
	@migrate -path migrations -database postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable up