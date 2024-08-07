package server

import (
	"context"
	"iosync/internal/services"
	"net/http"
	"time"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	context := context.Background()
	var request services.LoginRequest

	if err := s.ReadJson(w, r, &request); err != nil {
		s.ErrorJson(w, err)
		return
	}

	if err := ValidateInput(&request); err != nil {
		s.ErrorJson(w, err, http.StatusBadRequest)
		return
	}

	user, err := s.authService.AuthenticateUser(context, request)
	if err != nil {
		s.ErrorJson(w, err, http.StatusBadRequest)
		return
	}

	sessionId, err := s.sessionService.CreateSession(context, request.Username)
	if err != nil {
		s.ErrorJson(w, err, http.StatusBadRequest)
		return
	}
	s.SetCookie(w, "session_id", sessionId, time.Now().Add(30*time.Minute))

	response := Response{
		Message: "User Logged In!",
		Data:    user,
	}

	s.WriteJson(w, http.StatusOK, response)
}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	context := context.Background()
	var request services.RegisterRequest

	if err := s.ReadJson(w, r, &request); err != nil {
		s.ErrorJson(w, err)
		return
	}

	if err := ValidateInput(&request); err != nil {
		s.ErrorJson(w, err, http.StatusBadRequest)
		return
	}

	user, err := s.authService.AddUser(context, request)
	if err != nil {
		s.ErrorJson(w, err, http.StatusBadRequest)
		return
	}

	sessionId, err := s.sessionService.CreateSession(context, request.Username)
	if err != nil {
		s.ErrorJson(w, err, http.StatusBadRequest)
		return
	}
	s.SetCookie(w, "session_id", sessionId, time.Now().Add(30*time.Minute))

	response := Response{
		Message: "User Registered Successfully!",
		Data:    user,
	}

	s.WriteJson(w, http.StatusCreated, response)
}
