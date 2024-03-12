package routes

import "net/http"

type Router struct {
	Uri                  string
	Method               string
	Function             func(http.ResponseWriter, *http.Request)
	RequestAuthetication bool
}
