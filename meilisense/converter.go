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

var (
	MEILISEARCH_URL string
	MEILISEARCH_KEY string
	CSV_FILE_PATH   string
)

func init() {
	

	MEILISEARCH_URL = os.Getenv("MEILISEARCH_URL")
	MEILISEARCH_KEY = os.Getenv("MEILISEARCH_KEY")
	CSV_FILE_PATH = os.Getenv("CSV_FILE_PATH")

	if MEILISEARCH_URL == "" || MEILISEARCH_KEY == "" || CSV_FILE_PATH == "" {
		log.Fatal("Required environment variables are not set")
	}
}

func toInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Warning: Failed to convert '%s' to int. Defaulting to 0.", s)
		return 0
	}
	return val
}

func convertIndex(s string) int {
	val := s[1:]
	return toInt(val)
}

func toFloat(s string) float64 {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Printf("Warning: Failed to convert '%s' to float. Defaulting to 0.0.", s)
		return 0.0
	}
	return val
}

func InitClient() meilisearch.ServiceManager {
	client := meilisearch.New(MEILISEARCH_URL, meilisearch.WithAPIKey(MEILISEARCH_KEY))
	return client
}

func Convert() {
	file, err := os.Open(CSV_FILE_PATH)
	client := InitClient()

	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	reader.Read()
	if err != nil {
		log.Fatalf("Failed to read CSV header: %v", err)
	}

	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV rows: %v", err)
	}

	var documents []map[string]interface{}

	for _, row := range rows {
		doc := map[string]interface{}{
			"id":              toInt(row[0]),
			"title":           row[1],
			"score":           toFloat(row[2]),
			"rank":            convertIndex(row[3]),
			"studio":          row[4],
			"anime_type":      row[5],
			"episode":         toInt(row[6]),
			"description":     row[7],
			"rating":          row[8],
			"aired":           row[9],
			"image_url":       row[10],
			"genre":           strings.Split(row[11], ","),
			"recommended":     row[12],
			"mixed_feeling":   row[13],
			"not_recommended": row[14],
		}

		documents = append(documents, doc)
	}

	fmt.Printf("client: %v\n", client)

	if err != nil {
		log.Fatalf("Failed to initialize Meilisearch client: %v", err)
	}

	index := client.Index("anime")

	task, err := index.AddDocuments(documents)
	if err != nil {
		log.Fatalf("Failed to add documents to Meilisearch: %v", err)
	}

	fmt.Printf("Task ID: %v\n", task)
}

