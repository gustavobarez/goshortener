package handler

import (
	"goshortener/repository"
	"net/http"
	"strings"
)

// @BasePath /

// @Summary      Redirect to Long URL
// @Description  Finds a short URL by its ID and redirects the client to the original long URL.
// @Tags         URLs
// @Param        id   path      string  true  "The ID of the short URL"
// @Success      301  {string}  string  "Redirects to the original URL"
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /{id} [get]
func ShowUrlHandler(writer http.ResponseWriter, request *http.Request) {
	id := strings.TrimPrefix(request.URL.Path, "/")
	if id == "" {
		http.NotFound(writer, request)
		return
	}
	ctx := request.Context()

	urlItem, err := repository.FindById(ctx, id)
	if err != nil {
		sendError(writer, http.StatusInternalServerError, "internal server error")
		return
	}

	http.Redirect(writer, request, urlItem.OriginalURL, http.StatusMovedPermanently)
}
