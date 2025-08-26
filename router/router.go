package router

import (
	"net/http"
)

func Initialize() *http.ServeMux {

	mux := http.NewServeMux()
	initializeRoutes(mux)
	return mux

}
