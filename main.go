package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kasyap1234/anime-api/anime"
	"github.com/kasyap1234/anime-api/meilisense"
)

func main() {
	anime.InitDB()
	meilisense.Convert()
	r := chi.NewRouter()
	
	// r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/api/anime", anime.AnimeRouter())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":8080", r)
}
