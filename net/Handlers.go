package net

import (
	"encoding/json"
	"log"
	"net/http"

	"gitlab.com/archit-p/chitra/repo"
)

func baseDirHandler(w http.ResponseWriter, r *http.Request) {
	dir := "/home/archit-p/videos"
	movies := repo.GetMovieList(dir)

	moviesJson, err := json.Marshal(movies)

	if err != nil {
		log.Fatal("Failed to encode to JSON", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(moviesJson)
}

func movieDetailHandler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("v")

	if v == "" {
		log.Println("Received empty query for movie details")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	m := repo.GetMovieDetails(v)

	mJson, err := json.Marshal(m)

	if err != nil {
		log.Fatal("Failed to encode JSON", err)
	}

	w.Write(mJson)
}

func movieServeHandler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("v")

	if v == "" {
		log.Println("Received empty query for movie serving")
	}

	movie, err := repo.GetMoviePath(v)
	if err != nil {
		log.Println("Warning: ", err)
	}

	http.ServeFile(w, r, movie)
}

