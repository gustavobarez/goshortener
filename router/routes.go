package router

import (
	"goshortener/handler"
	"net/http"
)

func initializeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", func(writer http.ResponseWriter, router *http.Request) {
		if router.URL.Path == "/" && router.Method == http.MethodPost {
			handler.CreateUrlHandler(writer, router)
			return
		}
		if router.Method == http.MethodGet && router.URL.Path != "/" {
			handler.ShowUrlHandler(writer, router)
			return
		}
		http.NotFound(writer, router)
	})

}
