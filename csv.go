package main

import (
	"encoding/csv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

type Anime struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string
	Description string
	Genre       string
	ReleaseDate string
	Rating      float64
	Episodes    int
	Studio      string
	Status      string
	ImageURL    string
}

func read_csv() {
	file, err := os.Open("anime-2.csv")
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
     for _,record :=range records[1:]{
		anime :=Anime{
			
		}
	 }
}
