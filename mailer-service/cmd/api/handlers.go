package main

import (
	"log"
	"net/http"

	"github.com/tsawler/toolbox"
)

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload mailMessage

	err := app.Tools.ReadJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.Tools.ErrorJSON(w, err)
		return
	}

	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}

	err = app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		log.Println(err)
		app.Tools.ErrorJSON(w, err)
		return
	}

	payload := toolbox.JSONResponse{
		Error:   false,
		Message: "sent to " + requestPayload.To,
	}

	app.Tools.WriteJSON(w, http.StatusAccepted, payload)
}
