package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/healthcheck", app.ping)
	r.Post("/api/auth/register", app.registerUser)
	r.Post("/api/auth/login", app.loginUser)
	r.Post("/api/room", app.createRoom)
	r.Post("/api/room/join", app.joinRoom)
	r.Post("/api/ws/{room_id}", app.serveWebsocket)
	return r
}
