package handler

import (
	"encoding/json"
	"goshortener/repository"
	"goshortener/schemas"
	"math/rand"
	"net/http"
	"time"
)

// @BasePath /

// @Summary      Create Short URL
// @Description  Creates a new short URL from a provided long URL.
// @Tags         URLs
// @Accept       json
// @Produce      json
// @Param        request body handler.CreateUrlRequest true "Request body for creating a new short URL"
// @Success      201 {object} handler.CreateUrlResponse
// @Failure      400 {object} handler.ErrorResponse
// @Failure      500 {object} handler.ErrorResponse
// @Router       / [post]
func CreateUrlHandler(writer http.ResponseWriter, router *http.Request) {
	if router.Method != http.MethodPost {
		http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := router.Context()

	request := CreateURLRequest{}
	if err := json.NewDecoder(router.Body).Decode(&request); err != nil {
		http.Error(writer, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := request.Validate(); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	urlItem := schemas.URL{
		ID:          generateShortID(),
		OriginalURL: request.URL,
	}

	if err := repository.Save(ctx, urlItem); err != nil {
		logger.Errorf("failed to save")
		sendError(writer, http.StatusInternalServerError, "failed to create short url")
		return
	}

	sendSuccess(writer, http.StatusOK, "create-url", urlItem)

}

func generateShortID() string {
	const idLength = 6
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, idLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
