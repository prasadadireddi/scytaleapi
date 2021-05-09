package routes

import (
	"net/http"
	"github.com/gorilla/mux"

)

// Route struct
type Route struct {
	URI          string
	Method       string
	Handler      func(w http.ResponseWriter, r *http.Request)
}

// Load the routes
func Load() []Route {
	routes := basicRoutes
	routes = append(routes, svidRoutes...)
	return routes
}
func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {

			r.HandleFunc(route.URI,
				route.Handler,
			).Methods(route.Method)

	}
	return r
}
