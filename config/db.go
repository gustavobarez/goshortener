package config

import (
	"context"

	confaws "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func InitializeDynamoDB() (*dynamodb.Client, error) {
	logger = GetLogger("dynamodb")
	cfg, err := confaws.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Errorf("unable to load SDK config: %v", err)
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)
	return client, nil
}
