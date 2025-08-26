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
	tableName = "go-shortener-stack-GoshortenerTable-15NXOTCN9DYB7"
)

func Save(ctx context.Context, url schemas.URL) error {
	logger = config.GetLogger("repository")
	db = config.GetDynamoDB()
	av, err := attributevalue.MarshalMap(url)
	if err != nil {
		logger.Errorf("failed to marshal url item: %v", err)
		return err
	}

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

func FindById(ctx context.Context, id string) (schemas.URL, error) {
	logger = config.GetLogger("repository")
	db = config.GetDynamoDB()
	var urlItem schemas.URL
	key, err := attributevalue.MarshalMap(map[string]string{"id": id})
	if err != nil {
		logger.Errorf("failed to marshal key: %v", err)
		return urlItem, err
	}

	result, err := db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &tableName,
		Key: key,
	})

	if result.Item == nil {
		logger.Errorf("failed to find item: %v", err)
		return urlItem, err 
	}
	
	err = attributevalue.UnmarshalMap(result.Item, &urlItem)
	if err != nil {
		logger.Errorf("failed to unmarshal item: %v", err)
		return urlItem, err
	}
	return urlItem, nil
}