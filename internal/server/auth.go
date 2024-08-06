package server

import (
	"context"
	"iosync/internal/services"
	"net/http"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	var request services.LoginRequest

	if err := s.readJson(w, r, &request); err != nil {
		s.errorJson(w, err)
		return
	}

	if err := validateInput(&request); err != nil {
		s.errorJson(w, err, http.StatusBadRequest)
		return
	}

	if err := s.authService.AuthenticateUser(request); err != nil {
		s.errorJson(w, err, http.StatusBadRequest)
		return
	}

	response := Response{
		Message: "Login",
	}

	s.writeJson(w, http.StatusOK, response)
}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	context := context.Background()
	var request services.RegisterRequest

	if err := s.readJson(w, r, &request); err != nil {
		s.errorJson(w, err)
		return
	}

	if err := validateInput(&request); err != nil {
		s.errorJson(w, err, http.StatusBadRequest)
		return
	}

	user, err := s.authService.AddUser(context, request)
	if err != nil {
		s.errorJson(w, err, http.StatusBadRequest)
		return
	}

	response := Response{
		Message: "Success",
		Data:    user,
	}

	s.writeJson(w, http.StatusOK, response)
}
