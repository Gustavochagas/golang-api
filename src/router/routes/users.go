package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.SearchUsers,
		RequireAuthentication: true,
	}, {
		URI:                   "/users/{userId}",
		Method:                http.MethodGet,
		Function:              controllers.SearchUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/user/{userId}s",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{user_id}/follow",
		Method:                http.MethodPost,
		Function:              controllers.FollowUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{user_id}/unfollow",
		Method:                http.MethodPost,
		Function:              controllers.UnfollowUser,
		RequireAuthentication: true,
	},
}
