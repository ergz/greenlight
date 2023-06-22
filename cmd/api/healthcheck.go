package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	envelopedHealth := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, envelopedHealth, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "the server encountered an error and could process your request", http.StatusInternalServerError)
	}

}
