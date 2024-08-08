package server

import (
	"iosync/internal/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/sup", s.Sup)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", s.Login)
		r.Post("/register", s.Register)
	})

	r.Group(func(r chi.Router) {
		r.Use(middlewares.SessionMiddleware)

		r.Get("/protected", func(w http.ResponseWriter, r *http.Request) {
			username := r.Context().Value("session_id")
			s.WriteJson(w, 200, username)
		})
	})

	return r
}

func (s *Server) Sup(w http.ResponseWriter, r *http.Request) {
	responsePaylaod := Response{
		Message: "Sup",
	}

	s.WriteJson(w, http.StatusOK, responsePaylaod)
}
