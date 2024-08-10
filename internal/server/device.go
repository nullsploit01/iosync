package server

import (
	"net/http"
)

func (s *Server) AddDevice(w http.ResponseWriter, r *http.Request) {
	s.WriteJson(w, 200, "test")
}
