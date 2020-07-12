package net

import (
	"log"
	"net/http"
)

func Logger(inner http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		inner(w, r)
	}
}

