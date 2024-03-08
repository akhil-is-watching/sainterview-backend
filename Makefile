build:
	@go build -o bin/humanpal_service

run: build
	@./bin/humanpal_service
