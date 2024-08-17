package server

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (s *Server) CreateApiKey(w http.ResponseWriter, r *http.Request) {
	context := context.Background()
	idParam := chi.URLParam(r, "id")
	deviceId, err := strconv.Atoi(idParam)
	if err != nil {
		s.ErrorJson(w, errors.New("invalid device id"), http.StatusBadRequest)
		return
	}

	apiKey, err := s.apiKeyService.CreateApiKey(context, deviceId)
	if err != nil {
		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	responsPaylaod := Response{
		Data: apiKey,
	}

	s.WriteJson(w, http.StatusCreated, responsPaylaod)
}
