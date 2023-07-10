package routes

import (
	"api/src/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

// Route é um struct que representa todas as rotas da API
type Route struct {
	URI                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

func Configure(router *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, publishesRoutes...)

	for _, route := range routes {
		if route.RequireAuthentication {
			router.HandleFunc(
				route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}
	return router
}
