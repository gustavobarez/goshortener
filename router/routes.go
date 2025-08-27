package router

import (
	"goshortener/docs"
	"goshortener/handler"
	"net/http"
	"strings"

	httpSwagger "github.com/swaggo/http-swagger"
)

func initializeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(docs.SwaggerInfo.ReadDoc()))
	})

	mux.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
	))

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if strings.HasPrefix(request.URL.Path, "/swagger") {
			return
		}

		if request.URL.Path == "/" && request.Method == http.MethodPost {
			handler.CreateUrlHandler(writer, request)
			return
		}

		if request.Method == http.MethodGet && request.URL.Path != "/" {
			handler.ShowUrlHandler(writer, request)
			return
		}

		http.NotFound(writer, request)
	})
}
