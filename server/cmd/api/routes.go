package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.NotFound(app.notFound)
	mux.MethodNotAllowed(app.methodNotAllowed)

	mux.Use(app.logAccess)
	mux.Use(app.recoverPanic)
	mux.Use(middleware.RedirectSlashes)

	mux.Use(middleware.Heartbeat("/v1/ping"))

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Route("/v1", func(r chi.Router) {
		r.Get("/status", app.status)
		r.Route("/nodes", func(r chi.Router) {
			r.Get("/", app.GetNodes)
			r.Get("/{id}", app.GetNode)
			r.Get("/{nodeApiKey}/values", app.GetNodeValuesByAPIKey)
			r.Post("/{id}/apikey", app.GenerateNodeAPIKey)
			r.Post("/", app.CreateNode)
			r.Post("/value", app.AddNodeValue)
		})
	})

	return mux
}
