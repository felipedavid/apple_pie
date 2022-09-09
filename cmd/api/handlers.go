package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

func (app *application) createMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a new movie")
}

func (app *application) showMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "showing a movie")
}
