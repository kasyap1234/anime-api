package anime

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"strconv"
)

func getAllAnime(w http.ResponseWriter, r *http.Request) {

}

func getAnimeByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	// find anime by id ;
    for _,anime :=range animeList{
		if anime.ID== id{
			json.NewEncoder(w).Encode(anime); 
			return
		}
		http.Error(w,"Anime not found",http.StatusNotFound);
		return 
	}

}

func searchAnime(w http.ResponseWriter,r*http.Request){
	query :=r.URL.Query().Get("q")
	
}
