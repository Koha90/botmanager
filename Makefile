# build
build:
	@go build -o bin/botmanager cmd/main.go

run: build
	@./bin/botmanager

test:
	@go test ./...
