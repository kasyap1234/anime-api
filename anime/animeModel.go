package anime

type Anime struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Genre       string  `json:"genre"`
	ReleaseDate string  `json:"release_date"`
	Rating      float64 `json:"rating"`
	Episodes    int     `json:"episodes"`
	Studio      string  `json:"studio"`
	Status      string  `json:"status"`
	ImageURL    string  `json:"image_url"`
}
