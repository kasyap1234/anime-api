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


// searchResults, err := typesense.SearchAnime(client, "action")
// if err != nil {
//     fmt.Printf("Error searching: %v\n", err)
//     return
// }

// // Process the search results
// for _, hit := range searchResults.Hits {
//     fmt.Printf("Title: %s, Score: %f\n", hit.Document["Title"], hit.Document["Score"])
// }