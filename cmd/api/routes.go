package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheck)
	r.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovie)
	r.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovie)

	return r
}
