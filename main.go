package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kasyap1234/anime-api/anime"
	"github.com/kasyap1234/anime-api/typesense"
	"fmt"

)

func main() {
	anime.InitDB()
	client := typesense.NewClient("xyz", "http://localhost:8108")

	// Create the anime collection
	err := typesense.CreateAnimeCollection(client)
	if err != nil {
		fmt.Printf("Error creating collection: %v\n", err)
		return
	}

	// Convert CSV data to Typesense
	err = typesense.ConvertCSVToTypesense(client)
	if err != nil {
		fmt.Printf("Error converting CSV to Typesense: %v\n", err)
		return
	}

	fmt.Println("Anime data successfully indexed in Typesense!")
	searchResults, err := typesense.SearchAnime(client, "action", "score>7.0")
	if err != nil {
		fmt.Printf(err.Error())
	}
	for _, hit := range *searchResults.Hits {
		fmt.Printf("Title: %s, Score: %f\n", (*hit.Document)["Title"], (*hit.Document)["Score"])
	}
	r := chi.NewRouter()
	
	// r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/api/anime", anime.AnimeRouter())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":3000", r)

}