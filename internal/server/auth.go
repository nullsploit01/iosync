package server

import (
	"context"
	"errors"
	"fmt"
	"iosync/internal/services"
	"net/http"
	"time"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	context := context.Background()
	var request services.LoginRequest

	if err := s.readJson(w, r, &request); err != nil {
		s.errorJson(w, err)
		return
	}

	if err := validateInput(&request); err != nil {
		s.errorJson(w, err, http.StatusBadRequest)
		return
	}

	user, err := s.authService.AuthenticateUser(context, request)
	if err != nil {
		s.errorJson(w, err, http.StatusBadRequest)
		return
	}

	sessionId, err := s.sessionService.CreateSession(context, request.Username)
	if err != nil {
		s.errorJson(w, errors.New("error creating session, please try again"), http.StatusInternalServerError)
	}

	fmt.Println(sessionId)

	response := Response{
		Message: "Success",
		Data:    user,
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		Expires:  time.Now().Add(30 * time.Minute),
		HttpOnly: true, // Prevents access to the cookie via JavaScript
		Secure:   true, // Ensures the cookie is sent only over HTTPS
		Path:     "/",
		SameSite: http.SameSiteStrictMode, // Prevents CSRF attacks
	}

	http.SetCookie(w, cookie)

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

	s.writeJson(w, http.StatusCreated, response)
}
