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
	{
		URI:                   "/users/{user_id}/followers",
		Method:                http.MethodGet,
		Function:              controllers.SearchFollowers,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{user_id}/following",
		Method:                http.MethodGet,
		Function:              controllers.SearchFollowing,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{user_id}/update-password",
		Method:                http.MethodPost,
		Function:              controllers.UpdatePassword,
		RequireAuthentication: true,
	},
}
