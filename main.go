package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kasyap1234/anime-api/anime"
	_ "github.com/kasyap1234/anime-api/docs"
)
// @version 2.0
// @title Anime API
// @description This is an anime API server.
// @host localhost:3000
// @BasePath /api
func main() {
	anime.InitDB()
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/api/anime", anime.AnimeRouter())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":3000", r)
}
