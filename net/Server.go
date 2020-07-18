package net

import (
	"net/http"

	"github.com/gorilla/mux"
)

var ServerDir string

func StartAPIServer(dir string, port string) {
	ServerDir = dir
	router := NewRouter()

	http.ListenAndServe(":" + port, router)
}

func StartUIServer(dir string, port string) {
	r := mux.NewRouter()

	buildHandler := http.FileServer(http.Dir(dir))
	r.PathPrefix("/").Handler(buildHandler)

	http.ListenAndServe(":" + port, r)
}
