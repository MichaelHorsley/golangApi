package round

import (
	"poolgolang/src/infrastructure"
)

var controller = &Controller{Repository: Repository{}}

var routes = infrastructure.Routes{
	infrastructure.Route{
		"Index",
		"GET",
		"/rounds",
		controller.Index,
	},
	infrastructure.Route{
		"AddRound",
		"POST",
		"/rounds",
		controller.AddRound,
	},
	infrastructure.Route{
		"UpdateRound",
		"PUT",
		"/rounds/{id}",
		controller.UpdateRound,
	},
	infrastructure.Route{
		"DeleteRound",
		"DELETE",
		"/rounds/{id}",
		controller.DeleteRound,
	},
}

func GetRoutes() infrastructure.Routes {
	return routes
}
