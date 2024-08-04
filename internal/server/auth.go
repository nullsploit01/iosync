package server

import (
	"net/http"
)

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Login",
	}

	s.writeJson(w, http.StatusOK, response)
}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	var request RegisterRequest

	if err := s.readJson(w, r, &request); err != nil {
		s.errorJson(w, err)
		return
	}

	if err := validateInput(&request); err != nil {
		s.errorJson(w, err, http.StatusBadRequest)
		return
	}

	response := Response{
		Message: "Register",
	}

	s.writeJson(w, http.StatusOK, response)
}
