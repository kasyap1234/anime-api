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
	
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("anime.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	log.Printf("connecting to the database....")

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
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	var animeList []database.Anime
	var total int64

	db.Model(&database.Anime{}).Where("title LIKE ?", "%"+query+"%").Count(&total)
	result := db.Where("title LIKE ?", "%"+query+"%").Offset((page - 1) * pageSize).Limit(pageSize).Find(&animeList)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	response := PaginatedAnimeResponse{
		Data:       animeList,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int(math.Ceil(float64(total) / float64(pageSize))),
	}

	json.NewEncoder(w).Encode(response)
}

func searchAnimeByGenre(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("genre")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	var animeList []database.Anime
	var total int64

	db.Model(&database.Anime{}).Where("genre LIKE ?", "%"+query+"%").Count(&total)
	result := db.Where("genre LIKE ?", "%"+query+"%").Offset((page - 1) * pageSize).Limit(pageSize).Find(&animeList)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	response := PaginatedAnimeResponse{
		Data:       animeList,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int(math.Ceil(float64(total) / float64(pageSize))),
	}

	json.NewEncoder(w).Encode(response)
}

func searchAnimeByStudio(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("studio")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	var animeList []database.Anime
	var total int64

	db.Model(&database.Anime{}).Where("studio LIKE ?", "%"+query+"%").Count(&total)
	result := db.Where("studio LIKE ?", "%"+query+"%").Offset((page - 1) * pageSize).Limit(pageSize).Find(&animeList)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	response := PaginatedAnimeResponse{
		Data:       animeList,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int(math.Ceil(float64(total) / float64(pageSize))),
	}

	json.NewEncoder(w).Encode(response)
}

func searchAnimeByType(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("type")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	var animeList []database.Anime
	var total int64

	db.Model(&database.Anime{}).Where("anime_type LIKE ?", "%"+query+"%").Count(&total)
	result := db.Where("anime_type LIKE ?", "%"+query+"%").Offset((page - 1) * pageSize).Limit(pageSize).Find(&animeList)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	response := PaginatedAnimeResponse{
		Data:       animeList,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int(math.Ceil(float64(total) / float64(pageSize))),
	}

	json.NewEncoder(w).Encode(response)
}

func sortAnimeByScore(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	var animeList []database.Anime
	var total int64

	db.Model(&database.Anime{}).Count(&total)

	var result *gorm.DB
	if r.URL.Query().Get("sort") == "asc" {
		result = db.Order("score ASC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&animeList)
	} else {
		result = db.Order("score DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&animeList)
	}

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	response := PaginatedAnimeResponse{
		Data:       animeList,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int(math.Ceil(float64(total) / float64(pageSize))),
	}

	json.NewEncoder(w).Encode(response)
}
