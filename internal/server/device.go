package server

import (
	"context"
	"errors"
	"iosync/internal/services"
	"iosync/pkg/constants"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

	username, err := GetHttpRequestContextValue(r, constants.UsernameKey)

	if err != nil {
		s.ErrorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	requestBody.Username = username

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

func (s *Server) GetDevices(w http.ResponseWriter, r *http.Request) {
	context := context.Background()
	username, err := GetHttpRequestContextValue(r, constants.UsernameKey)

	if err != nil {
		s.ErrorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	devices, err := s.deviceService.GetDevices(context, username)
	if err != nil {
		s.ErrorJson(w, err)
	}

	responsePayload := Response{
		Data: devices,
	}

	s.WriteJson(w, http.StatusOK, responsePayload)
}

func (s *Server) GetDevice(w http.ResponseWriter, r *http.Request) {
	context := context.Background()
	idParam := chi.URLParam(r, "id")
	deviceId, err := strconv.Atoi(idParam)
	if err != nil {
		s.ErrorJson(w, errors.New("invalid device id"), http.StatusBadRequest)
	}

	device, err := s.deviceService.GetDevice(context, deviceId)
	if err != nil {
		s.ErrorJson(w, errors.New("device not found"), http.StatusBadRequest)
		return
	}

	responsePayload := Response{
		Data: device,
	}

	s.WriteJson(w, http.StatusOK, responsePayload)

}
