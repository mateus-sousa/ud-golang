package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.GetUsers,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodGet,
		Function:              controllers.GetUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/follow",
		Method:                http.MethodPost,
		Function:              controllers.FollowUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/stop-follow",
		Method:                http.MethodPost,
		Function:              controllers.StopFollowUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/followers",
		Method:                http.MethodGet,
		Function:              controllers.GetFollowers,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/following",
		Method:                http.MethodGet,
		Function:              controllers.GetFollowing,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/update-password",
		Method:                http.MethodPost,
		Function:              controllers.UpdatePassword,
		RequireAuthentication: true,
	},
}
