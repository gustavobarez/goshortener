package config

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	logger *Logger
	db     *dynamodb.Client
	once   sync.Once
)

func Init() {
	once.Do(func() {
		var err error
		db, err = InitializeDynamoDB(context.TODO())
		if err != nil {
			panic(fmt.Errorf("error initializing db: %v", err))
		}
	})
}

func GetDynamoDB() *dynamodb.Client {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}