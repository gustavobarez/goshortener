package config

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	logger *Logger
	db     *dynamodb.Client
)

func Init() error {
	var err error
	_, err = InitializeDynamoDB()
	if err != nil {
		return fmt.Errorf("error initializing db: %v", err)
	}
	return nil
}

func GetDynamoDB() *dynamodb.Client {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
