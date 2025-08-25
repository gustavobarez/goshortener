.PHONY: default run build test docs clean deploy
APP_NAME=goshortener

default: run-with-docs

run:
	@go run main.go

run-with-docs:
	@swag init
	@go run main.go

build:
	@swag init
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap main.go
build-GoshortenerFunction:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(ARTIFACTS_DIR)/bootstrap main.go
test:
	@go test ./ ...

docs:
	@swag init

deploy:
	sam deploy

clean:
	@rm -f $(APP_NAME)
	@rm -f bootstrap
	@rm -rf ./docs