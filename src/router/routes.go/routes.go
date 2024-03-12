package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Uri                  string
	Method               string
	Function             func(http.ResponseWriter, *http.Request)
	RequestAuthetication bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := RoutesUsers

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}

	return r
}
