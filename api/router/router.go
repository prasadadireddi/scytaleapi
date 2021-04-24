
package router

import (
"github.com/gorilla/mux"
"github.com/prasadadireddi/scytaleapi/api/router/routes"
)

// New routes

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutes(r)
}
