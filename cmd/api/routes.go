package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthCheck)
	mux.HandleFunc("/v1/movies/", app.movies)

	return mux
}
