package main

import (
	"logger-service/data"
	"net/http"

	"github.com/tsawler/toolbox"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// read json into var
	var requestPayload JSONPayload
	_ = app.Tools.ReadJSON(w, r, &requestPayload)

	// insert data
	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.Tools.ErrorJSON(w, err)
		return
	}

	resp := toolbox.JSONResponse{
		Error:   false,
		Message: "logged",
	}

	app.Tools.WriteJSON(w, http.StatusAccepted, resp)
}
