package config

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func InitializeDynamoDB(ctx context.Context) (*dynamodb.Client, error) {
	logger = GetLogger("dynamodb")
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Errorf("unable to load SDK config: %v", err)
		return nil, err
	}

	if endpoint := os.Getenv("AWS_ENDPOINT"); endpoint != "" {
		cfg.EndpointResolver = aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{URL: endpoint}, nil
		})
	}

	client := dynamodb.NewFromConfig(cfg)
	return client, nil
}
