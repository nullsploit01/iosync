package main

import (
	"net/http"

	"github.com/nullsploit01/iosync/internal/response"
)

func (app *application) status(w http.ResponseWriter, r *http.Request) {
	responsePayload := response.Response{
		Data: "sup",
	}

	err := response.JSON(w, http.StatusOK, responsePayload)
	if err != nil {
		app.serverError(w, r, err)
	}
}
