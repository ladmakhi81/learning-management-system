build-server:
	@go build -o ./bin/app ./cmd/api/main.go
run-server:build-server
	@./bin/app