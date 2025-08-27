package main

import (
	"context"
	"goshortener/config"
	"goshortener/router"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

var (
	logger  *config.Logger
	adapter *httpadapter.HandlerAdapter
)

func init() {
	config.Init()
	mux := router.Initialize()
	adapter = httpadapter.New(mux)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return adapter.ProxyWithContext(ctx, request)
}

func main() {
	logger = config.GetLogger("main")
	lambda.Start(Handler)
}
