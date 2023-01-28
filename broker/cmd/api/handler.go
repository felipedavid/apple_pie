package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *app) broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	js, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		_, _ = fmt.Fprintf(w, "Something bad happened... t-t")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	_, _ = w.Write(js)
}