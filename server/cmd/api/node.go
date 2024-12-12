package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/nullsploit01/iosync/internal/response"
)

func (app *application) GetNodes(w http.ResponseWriter, r *http.Request) {
	nodes, err := app.services.nodeService.GetNodes(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := map[string]any{
		"Data": nodes,
	}

	err = response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) GetNode(w http.ResponseWriter, r *http.Request) {
	nodeIdStr := r.URL.Query().Get("id")

	if nodeIdStr == "" {
		app.badRequest(w, r, fmt.Errorf("node id is required"))
		return
	}

	nodeId, err := strconv.Atoi(nodeIdStr)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	node, err := app.services.nodeService.GetNode(r.Context(), nodeId)
	data := map[string]any{
		"Data": node,
	}

	err = response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}
