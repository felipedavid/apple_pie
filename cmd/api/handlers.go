package main

import (
	"apple_pie/internal/data"
	"fmt"
	"net/http"
	"time"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorMessage(w, r, err)
	}
}

func (app *application) createGameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Creating a new movie")
}

func (app *application) showGameHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	data := data.Game{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Call of duty",
		Year:      2020,
		Genre:     []string{"Shooter"},
		Version:   0,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": data}, nil)
	if err != nil {
		app.serverErrorMessage(w, r, err)
	}
}
