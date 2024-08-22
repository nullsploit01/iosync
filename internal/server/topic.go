package server

import (
	"context"
	"errors"
	"iosync/ent"
	"iosync/internal/services"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (s *Server) CreateTopic(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	idParam := chi.URLParam(r, "id")
	deviceId, err := strconv.Atoi(idParam)
	if err != nil {
		s.ErrorJson(w, errors.New("invalid device id"), http.StatusBadRequest)
		return
	}

	var request services.CreateTopicPayload

	if err := s.ReadJson(w, r, &request); err != nil {
		s.ErrorJson(w, err)
		return
	}

	if err := ValidateInput(&request); err != nil {
		s.ErrorJson(w, err)
		return
	}

	topic, err := s.topicService.CreateTopic(ctx, deviceId, &request)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("device not found"), http.StatusBadRequest)
			return
		}

		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	responsePayload := &Response{
		Data: topic,
	}

	s.WriteJson(w, http.StatusCreated, responsePayload)
}
