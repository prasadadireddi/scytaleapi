

package routes

import (
"github.com/prasadadireddi/scytaleapi/api/controllers"
"net/http"
)

var svidRoutes = []Route{
	Route{
		URI:     "/api/v1/svid/validate/{spiffeid}",
		Method:  http.MethodGet,
		Handler: controllers.ValidateSpiffeID,
	},
}
