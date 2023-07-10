package routes

import (
	"api/src/controllers"
	"net/http"
)

var publishesRoutes = []Route{
	{
		URI:                   "/publishes",
		Method:                http.MethodPost,
		Function:              controllers.CreatePublish,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publishes",
		Method:                http.MethodGet,
		Function:              controllers.GetPublishes,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publishes/{publishId}",
		Method:                http.MethodGet,
		Function:              controllers.GetPublish,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publishes/{publishId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdatePublish,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publishes/{publishId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeletePublish,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/publishes",
		Method:                http.MethodGet,
		Function:              controllers.GetPublishesByUser,
		RequireAuthentication: true,
	},
}
