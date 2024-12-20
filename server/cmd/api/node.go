package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nullsploit01/iosync/internal/request"
	"github.com/nullsploit01/iosync/internal/response"
	"github.com/nullsploit01/iosync/internal/service"
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
	nodeIdStr := chi.URLParam(r, "id")

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

func (app *application) CreateNode(w http.ResponseWriter, r *http.Request) {
	var body service.CreateNodeRequest
	if err := request.DecodeJSON(w, r, &body); err != nil {
		app.badRequest(w, r, err)
		return
	}

	if err := request.ValidateRequest(&body); err != nil {
		app.badRequest(w, r, err)
		return
	}

	node, err := app.services.nodeService.CreateNode(r.Context(), body)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := map[string]any{
		"Data": node,
	}

	err = response.JSON(w, http.StatusCreated, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) AddNodeValue(w http.ResponseWriter, r *http.Request) {
	var body service.AddNodeValueRequest
	if err := request.DecodeJSON(w, r, &body); err != nil {
		app.badRequest(w, r, err)
		return
	}

	if err := request.ValidateRequest(&body); err != nil {
		app.badRequest(w, r, err)
		return
	}

	nodeValue, err := app.services.nodeService.AddNodeValue(r.Context(), body)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := map[string]any{
		"Data": nodeValue,
	}

	err = response.JSON(w, http.StatusCreated, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) GenerateNodeAPIKey(w http.ResponseWriter, r *http.Request) {
	nodeIdStr := chi.URLParam(r, "id")

	if nodeIdStr == "" {
		app.badRequest(w, r, fmt.Errorf("node id is required"))
		return
	}

	var body service.GenerateNodeAPIKeyRequest
	if err := request.DecodeJSON(w, r, &body); err != nil {
		app.badRequest(w, r, err)
		return
	}

	if err := request.ValidateRequest(&body); err != nil {
		app.badRequest(w, r, err)
		return
	}

	nodeId, err := strconv.Atoi(nodeIdStr)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	nodeApiKey, err := app.services.nodeService.GenerateNodeAPIKey(r.Context(), nodeId, body)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := map[string]any{
		"Data": nodeApiKey,
	}

	err = response.JSON(w, http.StatusCreated, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}
