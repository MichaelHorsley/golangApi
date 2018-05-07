package participant

import (
	"poolgolang/src/infrastructure"
)

var controller = &Controller{Repository: Repository{}}

var routes = infrastructure.Routes{
	infrastructure.Route{
		"Index",
		"GET",
		"/participants",
		controller.Index,
	},
	infrastructure.Route{
		"AddParticipant",
		"POST",
		"/participants",
		controller.AddParticipant,
	},
	infrastructure.Route{
		"UpdateParticipant",
		"PUT",
		"/participants/{id}",
		controller.UpdateParticipant,
	},
	infrastructure.Route{
		"DeleteParticipant",
		"DELETE",
		"/participants/{id}",
		controller.DeleteParticipant,
	},
}

func GetRoutes() infrastructure.Routes {
	return routes
}
