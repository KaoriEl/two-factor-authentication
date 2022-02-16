package routes

import (
	"github.com/gorilla/mux"
	"log"
	"main/internal/server/controllers"
	"net/http"
)

//Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes list
type Routes []Route

var routes = Routes{
	Route{
		"GivePublicKey",
		"POST",
		"/GivePublicKey",
		controllers.Index,
	},
	Route{
		"GivePublicKeyAdmin",
		"POST",
		"/GivePublicKeyAdmin",
		controllers.Index,
	},
}

//NewRouter func
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
