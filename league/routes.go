package league

import (
	"poolgolang/infrastructure"
)

var controller = &Controller{Repository: Repository{}}

var routes = infrastructure.Routes{
	infrastructure.Route{
		"Index",
		"GET",
		"/leagues",
		controller.Index,
	},
	infrastructure.Route{
		"AddLeague",
		"POST",
		"/leagues",
		controller.AddLeague,
	},
	infrastructure.Route{
		"UpdateLeague",
		"PUT",
		"/leagues/{id}",
		controller.UpdateLeague,
	},
	infrastructure.Route{
		"DeleteLeague",
		"DELETE",
		"/leagues/{id}",
		controller.DeleteLeague,
	},
}

func GetRoutes() infrastructure.Routes {
	return routes
}
