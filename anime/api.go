package anime

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kasyap1234/anime-api/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("anime.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	log.Printf("connecting to the database...."); 


	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&database.Anime{})
}

// PaginatedAnimeResponse represents the response structure for paginated anime data
type PaginatedAnimeResponse struct {
	Data       []database.Anime `json:"data"`
	Total      int64            `json:"total"`
	Page       int              `json:"page"`
	PageSize   int              `json:"pageSize"`
	TotalPages int              `json:"totalPages"`
}

// @Summary Get all anime
// @Description Get a list of all anime with optional pagination
// @Produce json
// @Param page query int false "Page number"
// @Param pageSize query int false "Number of items per page"
// @Success 200 {object} PaginatedAnimeResponse
// @Router /anime [get]
func getAllAnime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	var animes []database.Anime
	var total int64

	db.Model(&database.Anime{}).Count(&total)
	result := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&animes)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	response := PaginatedAnimeResponse{
		Data:       animes,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int(math.Ceil(float64(total) / float64(pageSize))),
	}

	json.NewEncoder(w).Encode(response)
}

// @Summary Get anime by ID
// @Description Get a single anime by its ID
// @Produce json
// @Param id path int true "Anime ID"
// @Success 200 {object} database.Anime
// @Router /anime/{id} [get]
func getAnimeByID(w http.ResponseWriter, r *http.Request) {
	var anime database.Anime
	id := chi.URLParam(r, "id")
	db.First(&anime, id)
	json.NewEncoder(w).Encode(anime)
}

// @Summary Search anime
// @Description Search for anime by name
// @Produce json
// @Param name query string true "Anime name"
// @Success 200 {array} database.Anime
// @Router /anime/search [get]
func searchAnime(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("name")
    var animeList []database.Anime
    result := db.Where("title LIKE ?", "%"+query+"%").Find(&animeList)

    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    // Filter results to include only continuous matches
    var filteredList []database.Anime
    for _, anime := range animeList {
        if strings.Contains(strings.ToLower(anime.Title), strings.ToLower(query)) {
            filteredList = append(filteredList, anime)
        }
    }

    json.NewEncoder(w).Encode(filteredList)
}

