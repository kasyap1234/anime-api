package anime

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	
)

func AnimeRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	

	r.Get("/", getAllAnime)
	r.Get("/{id}", getAnimeByID)
	r.Get("/search", searchAnime)

	return r
}
