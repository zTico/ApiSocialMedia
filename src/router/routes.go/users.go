package routes

import (
	"api/src/controllers"
	"net/http"
)

var RoutesUsers = []Router{
	{
		Uri:                  "/usuarios",
		Method:               http.MethodPost,
		Function:             controllers.Create,
		RequestAuthetication: false,
	},

	{
		Uri:                  "/usuarios",
		Method:               http.MethodGet,
		Function:             controllers.FindUsers,
		RequestAuthetication: false,
	},

	{
		Uri:                  "/usuarios/{userId}",
		Method:               http.MethodGet,
		Function:             controllers.FindUser,
		RequestAuthetication: false,
	},

	{
		Uri:                  "/usuarios/{userId}",
		Method:               http.MethodPut,
		Function:             controllers.Update,
		RequestAuthetication: false,
	},

	{
		Uri:                  "/usuarios/{userId}",
		Method:               http.MethodDelete,
		Function:             controllers.Delete,
		RequestAuthetication: false,
	},
}
