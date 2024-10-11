package typesense

import (
	"context"
	"github.com/AlekSi/pointer"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
)

func NewClient(apiKey string, host string) *typesense.Client {
	client := typesense.NewClient(
		typesense.WithAPIKey(apiKey),
		typesense.WithServer(host),
	)
	return client

}
func CreateAnimeCollection(client *typesense.Client) error {
	schema := &api.CollectionSchema{
		Name: "anime",
		Fields: []api.Field{
			{Name: "ID", Type: "string"},
			{Name: "Title", Type: "string"},
			{Name: "Score", Type: "float"},
			{Name: "Rank", Type: "int32"},
			{Name: "Studio", Type: "string"},
			{Name: "AnimeType", Type: "string"},
			{Name: "Episode", Type: "int32"},
			{Name: "Description", Type: "string"},
			{Name: "Rating", Type: "string"},
			{Name: "Aired", Type: "string"},
			{Name: "ImageUrl", Type: "string"},
			{Name: "Genre", Type: "string[]"},
			{Name: "recommended", Type: "int32"},
			{Name: "mixed_feeling", Type: "int32"},
			{Name: "not_recommended", Type: "int32"},
		},
		DefaultSortingField: pointer.ToString("Score"),
	}

	_, err := client.Collections().Create(context.Background(),schema)
	return err

}
