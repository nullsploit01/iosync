package server

import (
	"iosync/internal/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/sup", s.Sup)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", s.Login)
		r.Post("/register", s.Register)
	})

	r.Group(func(r chi.Router) {
		r.Use(middlewares.SessionMiddleware(s.sessionService))

		r.Get("/auth/me", s.GetSession)
		r.Post("/auth/logout", s.Logout)

		r.Route("/devices", func(r chi.Router) {
			r.Get("/", s.GetDevices)
			r.Get("/{id}", s.GetDevice)
			r.Delete("/{id}", s.DeleteDevice)
		})
	})

	return r
}

func (s *Server) Sup(w http.ResponseWriter, r *http.Request) {
	responsePayload := Response{
		Message: "Sup",
	}

	s.WriteJson(w, http.StatusOK, responsePayload)
}
