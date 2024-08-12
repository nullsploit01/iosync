package server

import (
	"context"
	"iosync/internal/services"
	"iosync/pkg/constants"
	"net/http"
	"time"
)

type AuthResponsePayload struct {
	SessionId string    `json:"sessionId"`
	Username  string    `json:"username"`
	ExpiresAt time.Time `json:"expiresAt"`
}

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

	_, err := s.authService.AuthenticateUser(context, request)
	if err != nil {
		s.ErrorJson(w, err, http.StatusBadRequest)
		return
	}

	session, err := s.sessionService.CreateSession(context, request.Username)
	if err != nil {
		s.ErrorJson(w, err, http.StatusBadRequest)
		return
	}

	s.SetCookie(w, string(constants.SessionIDCookieKey), session.SessionID, time.Now().Add(30*time.Minute))

	response := Response{
		Message: "User Logged In!",
		Data: AuthResponsePayload{
			SessionId: session.SessionID,
			Username:  session.Username,
			ExpiresAt: session.ExpiresAt,
		},
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

	_, err := s.authService.AddUser(context, request)
	if err != nil {
		s.ErrorJson(w, err, http.StatusBadRequest)
		return
	}

	session, err := s.sessionService.CreateSession(context, request.Username)
	if err != nil {
		s.ErrorJson(w, err, http.StatusBadRequest)
		return
	}
	s.SetCookie(w, string(constants.SessionIDCookieKey), session.SessionID, time.Now().Add(30*time.Minute))

	response := Response{
		Message: "User Registered Successfully!",
		Data: AuthResponsePayload{
			SessionId: session.SessionID,
			Username:  session.Username,
			ExpiresAt: session.ExpiresAt,
		},
	}

	s.WriteJson(w, http.StatusCreated, response)
}
