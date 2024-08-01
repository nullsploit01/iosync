package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/sup", s.Sup)

	return r
}

func (s *Server) Sup(w http.ResponseWriter, r *http.Request) {
	responsePaylaod := Response{
		Message: "Sup",
	}

	s.writeJson(w, http.StatusOK, responsePaylaod)
}
