package net

import (
	"encoding/json"
	"log"
	"net/http"

	"gitlab.com/archit-p/chitra/repo"
)

func baseDirHandler(w http.ResponseWriter, r *http.Request) {
	movies, err := repo.GetMovieList(ServerDir)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		log.Printf("No files exist at %s", ServerDir)
		w.WriteHeader(http.StatusServiceUnavailable)

		statusMessage := "Requested directory is empty."

		if err := json.NewEncoder(w).Encode(statusMessage); err != nil {
			panic(err)
		}
		return
	}

	moviesJson, err := json.Marshal(movies)

	if err != nil {
		log.Printf("baseDirHandler: %s", err)
		w.WriteHeader(http.StatusInternalServerError)

		statusMessage := "JSON encoder failed"
		if err := json.NewEncoder(w).Encode(statusMessage); err != nil {
			panic(err)
		}
		return
	}

	w.Write(moviesJson)
}

func movieDetailHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("v")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if key == "" {
		log.Println("movieDetailHandler: Received empty query for movie details")
		w.WriteHeader(http.StatusBadGateway)

		statusMessage := "Didn't receive a query"

		if err := json.NewEncoder(w).Encode(statusMessage); err != nil {
			panic(err)
		}
		return
	}

	movieDetail, err := repo.GetMovieDetails(key)

	if err != nil {
		log.Printf("movieDetailHandler: No movie exists with %s", key)
		w.WriteHeader(http.StatusServiceUnavailable)

		statusMessage := "Received query is illegal"
		if err := json.NewEncoder(w).Encode(statusMessage); err != nil {
			panic(err)
		}
		return
	}

	detailJson, err := json.Marshal(movieDetail)

	if err != nil {
		log.Printf("movieDetailHandler: %s", err)
		w.WriteHeader(http.StatusInternalServerError)

		statusMessage := "JSON encoder failed"
		if err := json.NewEncoder(w).Encode(statusMessage); err != nil {
			panic(err)
		}
		return
	}

	w.Write(detailJson)
}

func movieServeHandler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("v")

	if v == "" {
		log.Println("Received empty query for movie serving")
	}

	movie, err := repo.GetMoviePath(v)
	if err != nil {
		log.Println("movieServeHandler: ", err)
	}

	http.ServeFile(w, r, movie)
}
