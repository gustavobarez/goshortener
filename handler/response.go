package handler

import (
	"encoding/json"
	"fmt"
	"goshortener/schemas"
	"net/http"
)

func sendError(writer http.ResponseWriter, code int, msg string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(ErrorResponse{
		Message:   msg,
		ErrorCode: http.StatusText(code),
	})
}

func sendSuccess(writer http.ResponseWriter, code int, op string, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(map[string]interface{}{
		"message": fmt.Sprintf("operation from handler: %s successful", op),
		"data":    data,
	})
}

type URLResponse struct {
	ID          string `json:"id"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateUrlResponse struct {
	Message string      `json:"message"`
	Data    schemas.URL `json:"data"`
}
