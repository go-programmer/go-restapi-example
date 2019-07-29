package users

import "github.com/codeanit/go-practice/utilities"

var UsersRoutes = utilities.Routes{
	utilities.Route{
		Name:        "GetUser",
		Method:      "GET",
		Pattern:     "/user/{userId}",
		HandlerFunc: GetUser,
	},
	utilities.Route{
		Name:        "PostUser",
		Method:      "POST",
		Pattern:     "/user",
		HandlerFunc: PostUser,
	},
	utilities.Route{
		Name:        "DeleteUser",
		Method:      "GET",
		Pattern:     "/user/{userId}",
		HandlerFunc: DeleteUser,
	},
}
