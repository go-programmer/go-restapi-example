package user

import Util "mywebapp/v9/utilities"

var UserRoutes = Util.Routes{
	// Util.Route{
	// 	"GetUser",
	// 	"GET",
	// 	"/user",
	// 	GetUser,
	// },
	Util.Route{
		"PostUser",
		"POST",
		"/user",
		PostUser,
	},
	// Util.Route{
	// 	"DeleteUser",
	// 	"GET",
	// 	"/user/{todoId}",
	// 	DeleteUser,
	// },
}
