package match

import (
	"poolgolang/src/infrastructure"
)

var controller = &Controller{Repository: Repository{}}

var routes = infrastructure.Routes{
	infrastructure.Route{
		"Index",
		"GET",
		"/matches",
		controller.Index,
	},
	infrastructure.Route{
		"AddMatch",
		"POST",
		"/matches",
		controller.AddMatch,
	},
	infrastructure.Route{
		"UpdateMatch",
		"PUT",
		"/matches/{id}",
		controller.UpdateMatch,
	},
	infrastructure.Route{
		"DeleteMatch",
		"DELETE",
		"/matches/{id}",
		controller.DeleteMatch,
	},
}

func GetRoutes() infrastructure.Routes {
	return routes
}
