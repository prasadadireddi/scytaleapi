

package routes

import (
"github.com/prasadadireddi/scytaleapi/api/controllers"
"net/http"
)

var basicRoutes = []Route{
	Route{
		URI:     "/api/v1/ping",
		Method:  http.MethodGet,
		Handler: controllers.Ping,
	},
	Route{
		URI:     "/api/v1/workloads",
		Method:  http.MethodGet,
		Handler: controllers.GetWorkloads,
	},
	Route{
		URI:     "/api/v1/workload/{selector}",
		Method:  http.MethodGet,
		Handler: controllers.GetWorkloadsBySelector,
	},
	Route{
		URI:     "/api/v1/workload",
		Method:  http.MethodPost,
		Handler: controllers.CreateWorkload,
	},
	Route{
		URI:     "/api/v1/workload/{spiffeid}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateWorkload,
	},
	Route{
		URI:     "/api/v1/workload/{spiffeid}/{selector}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateSelector,
	},
	Route{
		URI:     "/api/v1/workload/{spiffeid}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteWorkload,
	},
	Route{
		URI:     "/api/v1/workload/{spiffeid}/{selector}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteSelector,
	},
}
