package router

import (
	"api/src/router/routes.go"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.Configure(r)
}
