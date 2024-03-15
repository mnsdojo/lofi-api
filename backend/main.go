package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/mnsdojo/lofi-api/backend/handlers"
	"github.com/mnsdojo/lofi-api/backend/repository"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	songRepo := &repository.InMemorySongRepo{}
	songHandler := handlers.NewSongHandler(songRepo)

	r.Get("/", songHandler.GetSongs)

	http.ListenAndServe(":3000", r)
}
