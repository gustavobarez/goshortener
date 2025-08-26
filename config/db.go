package config

import (
	"context"

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

	client := dynamodb.NewFromConfig(cfg)
	return client, nil
}
