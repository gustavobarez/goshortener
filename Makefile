.PHONY: default run build test docs clean deploy
APP_NAME=goshortener

default: run-with-docs

run:
	@sam local start-api --env-vars env.json --profile default --region us-east-1
run-with-docs:
	@swag init
	@sam local start-api --env-vars env.json --profile default --region us-east-1
build:
	@swag init
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(ARTIFACTS_DIR)/bootstrap main.go
build-GoshortenerFunction:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(ARTIFACTS_DIR)/bootstrap main.go
test:
	@go test ./ ...
docs:
	@swag init
deploy:
	@sam deploy --guided
clean:
	@rm -f $(APP_NAME)
	@rm -f bootstrap
	@rm -rf ./docs