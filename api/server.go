package api


import (
"fmt"
"github.com/gorilla/handlers"
"github.com/gorilla/mux"
"log"
"net/http"

"github.com/prasadadireddi/scytaleapi/auto"
"github.com/prasadadireddi/scytaleapi/config"

"github.com/prasadadireddi/scytaleapi/api/router"
)

func init() {
	config.Load()
	auto.Load()
}

// Run message
func Run() {
	fmt.Printf("\nListening [::]:%d \n\n", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	corsMw := mux.CORSMethodMiddleware(r)
	r.Use(corsMw)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
