package server

import (
	"context"
	"errors"
	"fmt"
	"iosync/ent"
	"iosync/internal/services"
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

	var request services.CreateApiKeyRequest

	if err := s.ReadJson(w, r, &request); err != nil {
		s.ErrorJson(w, errors.New("invalid request body"))
		return
	}

	apiKey, err := s.apiKeyService.CreateApiKey(context, deviceId, &request)

	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("device not found"))
			return
		}

		var constraintError *ent.ConstraintError
		if errors.As(err, &constraintError) {
			s.ErrorJson(w, fmt.Errorf("device with ID %d already has an associated API key", deviceId))
			return
		}

		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	responsPaylaod := Response{
		Data: apiKey,
	}

	s.WriteJson(w, http.StatusCreated, responsPaylaod)
}

func (s *Server) GetApiKey(w http.ResponseWriter, r *http.Request) {
	context := context.Background()
	key := chi.URLParam(r, "key")
	if key == "" {
		s.ErrorJson(w, errors.New("invalid api key"), http.StatusBadRequest)
		return
	}

	idParam := chi.URLParam(r, "id")
	deviceId, err := strconv.Atoi(idParam)
	if err != nil {
		s.ErrorJson(w, errors.New("invalid device id"), http.StatusBadRequest)
		return
	}

	apiKey, err := s.apiKeyService.GetApiKey(context, deviceId, key)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("device or API key not found"), http.StatusBadRequest)
			return
		}

		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	response := Response{
		Data: apiKey,
	}

	s.WriteJson(w, http.StatusOK, response)
}

func (s *Server) RevokeApiKey(w http.ResponseWriter, r *http.Request) {
	context := context.Background()

	key := chi.URLParam(r, "key")
	if key == "" {
		s.ErrorJson(w, errors.New("invalid api key"), http.StatusBadRequest)
		return
	}

	idParam := chi.URLParam(r, "id")
	deviceId, err := strconv.Atoi(idParam)
	if err != nil {
		s.ErrorJson(w, errors.New("invalid device id"), http.StatusBadRequest)
		return
	}

	err = s.apiKeyService.RevokeApiKey(context, deviceId, key)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("device or API key not found"), http.StatusBadRequest)
			return
		}

		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	s.WriteJson(w, http.StatusOK, Response{})
}

func (s *Server) ActivateApiKey(w http.ResponseWriter, r *http.Request) {
	context := context.Background()

	key := chi.URLParam(r, "key")
	if key == "" {
		s.ErrorJson(w, errors.New("invalid api key"), http.StatusBadRequest)
		return
	}

	idParam := chi.URLParam(r, "id")
	deviceId, err := strconv.Atoi(idParam)
	if err != nil {
		s.ErrorJson(w, errors.New("invalid device id"), http.StatusBadRequest)
		return
	}

	err = s.apiKeyService.ActivateApiKey(context, deviceId, key)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("device or API key not found"), http.StatusBadRequest)
			return
		}

		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	s.WriteJson(w, http.StatusOK, Response{})
}
