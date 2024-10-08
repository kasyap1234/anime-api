package database 

import (
	"encoding/csv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

type Anime struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Score       float64
	Rank        uint
	Studio      string
	AnimeType   string
	Episode     uint
	Description string
	Rating      string
	Aired       string
	ImageUrl    string
	Genre       string
}

// ID,Title,Score,Rank,Studio,AnimeType,Episode,Description,Rating,Aired,ImageUrl,Genre,
func ReadCSV() {
	file, err := os.Open("anime-2-formatted.csv")
	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open(sqlite.Open("anime.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Anime{})
	for _, record := range records[1:] {
		id, err := strconv.ParseUint(record[0], 10, 32)
		score, err := strconv.ParseFloat(record[2], 64)
		// rank, err := strconv.ParseUint(record[3], 10, 64)
		rank :=ParseRank(record[3]); 

		if err != nil {
			log.Printf("Error parsing ID: %v", err)
			continue
		}
		anime := Anime{
			ID:          uint(id),
			Title:       record[1],
			Score:       score,
			Rank:        uint(rank),
			Studio:      record[4],
			AnimeType:   record[5],
			Episode:     uint(parseIntOrZero(record[6])),
			Description: record[7],
			Rating:      record[8],
			Aired:       record[9],
			ImageUrl:    record[10],
			Genre:       record[11],
		}
		db.Create(&anime)

	}
	log.Println("CSV file read successfully")

}
func ParseRank(s string) int {
	rank :=s[1:]; 
	rankInt,err :=strconv.Atoi(rank); 
	if err !=nil {
		return 0; 
	}
	return rankInt; 

}
func parseIntOrZero(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i

}

func parseFloatOrZero(s string) float64 {
	float, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return float

}
