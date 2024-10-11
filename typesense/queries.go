package typesense

import (
	"context"

	"github.com/AlekSi/pointer"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
)

func SearchAnime(client *typesense.Client, query string) (*api.SearchResult, error) {
	searchParameters := &api.SearchCollectionParams{
		Q:        query,
		QueryBy:  "Title,Description,Genre,Studio",
		FilterBy: pointer.ToString("Score:>=7"),
		SortBy:   pointer.ToString("Score:desc"),
		Limit:    pointer.ToInt(20),
	}

	return client.Collection("anime").Documents().Search(context.Background(), searchParameters)
}
