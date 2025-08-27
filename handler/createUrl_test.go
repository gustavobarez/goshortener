package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"goshortener/config"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var (
	db *dynamodb.Client
)

func createTestTable(ctx context.Context, db *dynamodb.Client, tableName string) error {
	_, err := db.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: &tableName,
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	})
	return err
}

func TestCreateUrlHandler(t *testing.T) {
	os.Setenv("AWS_ENDPOINT", "http://localhost:8000")
	os.Setenv("AWS_ACCESS_KEY_ID", "fake")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "fake")
	os.Setenv("AWS_REGION", "us-east-1")
	os.Setenv("DYNAMODB_TABLE", "urls-table-test")

	ctx := context.Background()
	db = config.GetDynamoDB()
	if db == nil {
		t.Fatal("dynamoDB client is nil")
	}
	tableName := "urls-table-test"
	os.Setenv("DYNAMODB_TABLE", tableName)

	createTestTable(ctx, db, tableName)

	t.Run("Should create URL successfully", func(t *testing.T) {
		payload := CreateURLRequest{
			URL: "https://www.google.com",
		}

		jsonPayload, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonPayload))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		CreateUrlHandler(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		var response CreateUrlResponse
		json.Unmarshal(w.Body.Bytes(), &response)

		if response.Data.ID == "" {
			t.Error("expected id to be generated")
		}

		if response.Data.OriginalURL != payload.URL {
			t.Errorf("expected url %s, got %s", payload.URL, response.Data.OriginalURL)
		}
	})

	t.Run("Should return error for empty URL", func(t *testing.T) {
		payload := CreateURLRequest{
			URL: "",
		}
		jsonPayload, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonPayload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		CreateUrlHandler(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", w.Code)
		}
	})

	t.Run("Should return error for invalid JSON", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		CreateUrlHandler(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", w.Code)
		}
	})

	t.Run("Should return error for invalid URL format", func(t *testing.T) {
		payload := CreateURLRequest{
			URL: "not-a-valid-url",
		}
		jsonPayload, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonPayload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		CreateUrlHandler(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", w.Code)
		}
	})

	t.Run("Should return error for wrong HTTP method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		CreateUrlHandler(w, req)

		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, got %d", w.Code)
		}
	})

}
