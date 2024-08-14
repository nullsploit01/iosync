package server

import (
	"context"
	"errors"
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

	sessionTimeOut := time.Now().Add(30 * time.Minute)
	if request.RememberMe {
		sessionTimeOut = time.Now().Add((720 * time.Hour)) // 30 Days
	}

	session, err := s.sessionService.CreateSession(context, request.Username, sessionTimeOut)
	if err != nil {
		s.ErrorJson(w, err, http.StatusBadRequest)
		return
	}

	s.SetCookie(w, string(constants.SessionIDKey), session.SessionID, sessionTimeOut)

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

	sessionTimeOut := time.Now().Add(30 * time.Minute)
	session, err := s.sessionService.CreateSession(context, request.Username, sessionTimeOut)
	if err != nil {
		s.ErrorJson(w, err, http.StatusBadRequest)
		return
	}
	s.SetCookie(w, string(constants.SessionIDKey), session.SessionID, time.Now().Add(30*time.Minute))

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

func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {
	context := context.Background()
	response := Response{
		Message: "success",
	}

	sessionId, err := GetHttpRequestContextValue(r, constants.SessionIDKey)
	if err != nil {
		s.WriteJson(w, http.StatusOK, response)
		return
	}

	err = s.sessionService.RemoveSession(context, sessionId)
	if err != nil {
		s.ErrorJson(w, errors.New("something went wrong"), http.StatusInternalServerError)
		return
	}

	s.WriteJson(w, http.StatusOK, response)
}

func (s *Server) GetSession(w http.ResponseWriter, r *http.Request) {
	context := context.Background()
	username, err := GetHttpRequestContextValue(r, constants.UsernameKey)
	if err != nil {
		s.ErrorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	session, err := s.sessionService.GetUserActiveSession(context, username)
	if err != nil {
		s.ErrorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	response := Response{
		Data: AuthResponsePayload{
			SessionId: session.SessionID,
			Username:  session.Username,
			ExpiresAt: session.ExpiresAt,
		},
	}

	s.WriteJson(w, http.StatusOK, response)
}
