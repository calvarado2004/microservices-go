package main

import (
	"calvarado2004/microservices-go/log-service/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	//read json into a var
	var requestPayload JSONPayload
	_ = app.readJSON(w, r, &requestPayload)

	//write to mongo
	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "Log entry created successfully",
	}

	app.writeJSON(w, http.StatusAccepted, resp)

}
