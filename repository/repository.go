package repository

import (
	"context"
	"goshortener/config"
	"goshortener/schemas"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	logger *config.Logger
	db     *dynamodb.Client
)

func Save(ctx context.Context, url schemas.URL) error {
	logger = config.GetLogger("repository")
	db = config.GetDynamoDB()
	av, err := attributevalue.MarshalMap(url)
	if err != nil {
		logger.Errorf("failed to marshal url item: %v", err)
		return err
	}

	tableName := "go-shortener-stack-GoshortenerTable-15NXOTCN9DYB7"
	_, err = db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &tableName,
		Item:      av,
	})
	if err != nil {
		logger.Errorf("failed to save url to dynamodb: %v", err)
		return err
	}
	return nil
}
