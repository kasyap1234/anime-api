package typesense

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/typesense/typesense-go/typesense"
	"os"
	"strconv"
	"strings"
)

func ConvertCSVToTypesense(client *typesense.Client) error {
	file, err := os.Open("anime-2-formatted.csv")
	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	// Index data into Typesense
	for _, record := range records[1:] { // Skip header row
		document := map[string]interface{}{
			"ID":              record[0],
			"Title":           record[1],
			"Score":           parseFloat(record[2]),
			"Rank":            parseInt(record[3]),
			"Studio":          record[4],
			"AnimeType":       record[5],
			"Episode":         parseInt(record[6]),
			"Description":     record[7],
			"Rating":          record[8],
			"Aired":           record[9],
			"ImageUrl":        record[10],
			"Genre":           strings.Split(record[11], ","),
			"recommended":     parseInt(record[12]),
			"mixed_feeling":   parseInt(record[13]),
			"not_recommended": parseInt(record[14]),
		}

		_, err := client.Collection("anime").Documents().Create(context.Background(), document)
		if err != nil {
			fmt.Printf("Error indexing document: %v\n", err)
		}
	}


	fmt.Println("Indexing complete!")
return err; 

}

func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
