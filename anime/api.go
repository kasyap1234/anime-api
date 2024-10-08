package anime

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kasyap1234/anime-api/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("anime.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&database.Anime{})

}
func animeList() []database.Anime {
	var animeList []database.Anime
	result := db.Find(&animeList)
	if result.Error != nil {
		panic("failed to fetch anime list")
	}
	return animeList

}

func getAllAnime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	animeList := animeList()

	json.NewEncoder(w).Encode(animeList)

}

func getAnimeByID(w http.ResponseWriter, r *http.Request) {
	var anime database.Anime
	id := chi.URLParam(r, "id")
	db.First(&anime, id)
	json.NewEncoder(w).Encode(anime)
}
