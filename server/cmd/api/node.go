package main

import (
	"net/http"

	"github.com/nullsploit01/iosync/internal/response"
)

func (app *application) GetNodes(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
	data := map[string]string{
		"Data": "nodes",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) GetNode(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
	data := map[string]string{
		"Data": "node",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}
