package net

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name		string
	Method		string
	Pattern		string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.HandlerFunc

		handler = route.HandlerFunc
		handler = Logger(handler)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes {
	Route{
		"Home",
		"GET",
		"/",
		baseDirHandler,
	},
	Route{
		"Watch",
		"GET",
		"/watch",
		movieDetailHandler,
	},
	Route{
		"Video",
		"GET",
		"/vid",
		movieServeHandler,
	},
}
