package server

import (
	"context"
	"iosync/internal/services"
	"iosync/pkg/constants"
	"net/http"
)

func (s *Server) AddDevice(w http.ResponseWriter, r *http.Request) {
	context := context.Background()
	var requestBody services.AddDeviceRequest

	if err := s.ReadJson(w, r, &requestBody); err != nil {
		s.ErrorJson(w, err)
		return
	}

	if err := ValidateInput(&requestBody); err != nil {
		s.ErrorJson(w, err, http.StatusBadRequest)
		return
	}

	username := r.Context().Value(constants.UsernameKey)

	if usernameStr, ok := username.(string); !ok {
		s.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	} else {
		requestBody.Username = usernameStr
	}

	device, err := s.deviceService.AddDevice(context, requestBody)
	if err != nil {
		s.WriteJson(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	responsePayload := Response{
		Message: "Device added successfully",
		Data:    device,
	}

	s.WriteJson(w, http.StatusCreated, responsePayload)
}
