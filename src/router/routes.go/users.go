package routes

import "net/http"

var RoutesUsers = []Router{
	{
		Uri:    "/usuarios",
		Method: http.MethodPost,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequestAuthetication: false,
	},

	{
		Uri:    "/usuarios",
		Method: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequestAuthetication: false,
	},

	{
		Uri:    "/usuarios/{userId}",
		Method: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequestAuthetication: false,
	},

	{
		Uri:    "/usuarios/{userId}",
		Method: http.MethodPut,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequestAuthetication: false,
	},

	{
		Uri:    "/usuarios/{userId}",
		Method: http.MethodDelete,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequestAuthetication: false,
	},
}
