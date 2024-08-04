package server

import (
	"net/http"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Login",
	}

	s.writeJson(w, http.StatusOK, response)
}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Register",
	}

	s.writeJson(w, http.StatusOK, response)
}
