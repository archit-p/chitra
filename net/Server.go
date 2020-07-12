package net

import (
	"net/http"

	"github.com/gorilla/mux"
)

func StartAPIServer() {
	router := NewRouter()

	http.ListenAndServe(":8080", router)
}

func StartUIServer() {
	r := mux.NewRouter()

	buildHandler := http.FileServer(http.Dir("ui/build"))
	r.PathPrefix("/").Handler(buildHandler)

	http.ListenAndServe(":5000", r)
}
