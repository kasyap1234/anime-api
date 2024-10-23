package meilisense


import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/meilisearch/meilisearch-go"
)

// Helper function to convert a string to int safely
func toInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Warning: Failed to convert '%s' to int. Defaulting to 0.", s)
		return 0
	}
	return val
}
func convertIndex(s string)int {
	val :=s[1:]; 
	

	return toInt(val);
}




// Helper function to convert a string to float safely
func toFloat(s string) float64 {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Printf("Warning: Failed to convert '%s' to float. Defaulting to 0.0.", s)
		return 0.0
	}
	return val
}


func InitClient()meilisearch.ServiceManager {

	client :=meilisearch.New("http://meilisearch:7700",meilisearch.WithAPIKey("masterKey"))

	return client

}
func Convert() {
	// Open the CSV file
	file, err := os.Open("anime-2-formatted.csv")
	client := InitClient(); 
	
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	// Initialize CSV reader
	reader := csv.NewReader(file)

	// Read the CSV header (first row)
	reader.Read()
	if err != nil {
		log.Fatalf("Failed to read CSV header: %v", err)

	}

	// Read the rest of the rows
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV rows: %v", err)
	}

	// Create a slice of maps to hold the documents
	var documents []map[string]interface{}

	// Iterate over each row and convert it to a map
	for _, row := range rows {
		doc := map[string]interface{}{
			"id":             toInt(row[0]),
			"title":          row[1],
			"score":          toFloat(row[2]),
			"rank":           convertIndex(row[3]),
			"studio":         row[4],
			"anime_type":     row[5],
			"episode":        toInt(row[6]),
			"description":    row[7],
			"rating":         row[8],
			"aired":          row[9],
			"image_url":      row[10],
			"genre":          strings.Split(row[11], ","), // Split genres by comma
			"recommended":    row[12],
			"mixed_feeling":  row[13],
			"not_recommended": row[14],
		}

		// Add the document to the slice
		documents = append(documents, doc)
	}

	// Initialize Meilisearch client
	
fmt.Printf("client: %v\n", client);

	if err != nil {
		log.Fatalf("Failed to initialize Meilisearch client: %v", err)
	}

	// Create or select the index (e.g., "anime")
	index := client.Index("anime")

	// Add documents to Meilisearch
	task, err := index.AddDocuments(documents)
	if err != nil {
		log.Fatalf("Failed to add documents to Meilisearch: %v", err)
	}

	fmt.Printf("Task ID: %v\n", task)

}

