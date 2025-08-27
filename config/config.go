package config

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
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
	if db != nil {
		return db
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				if service == dynamodb.ServiceID {
					return aws.Endpoint{
						URL:           "http://localhost:8000",
						SigningRegion: "us-east-1",
					}, nil
				}
				return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
			}),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	db = dynamodb.NewFromConfig(cfg)
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
