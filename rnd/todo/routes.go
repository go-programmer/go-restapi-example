package todo

import Util "mywebapp/v9/utilities"

var TodoRoutes = Util.Routes{
	Util.Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Util.Route{
		"TodoIndex",
		"GET",
		"/todo",
		TodoIndex,
	},
	Util.Route{
		"TodoCreate",
		"POST",
		"/todo",
		TodoCreate,
	},
	Util.Route{
		"TodoShow",
		"GET",
		"/todo/{todoId}",
		TodoShow,
	},
}
