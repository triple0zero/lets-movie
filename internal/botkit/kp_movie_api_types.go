package botkit

import "time"

type KinopoiskMovieId struct {
	ID         int `json:"id"`
	ExternalID struct {
		KpHD string `json:"kpHD"`
		Imdb string `json:"imdb"`
		Tmdb int    `json:"tmdb"`
	} `json:"externalId"`
	Name            string `json:"name"`
	AlternativeName string `json:"alternativeName"`
	EnName          string `json:"enName"`
	Names           []struct {
		Name     string `json:"name"`
		Language string `json:"language"`
		Type     string `json:"type"`
	} `json:"names"`
	Type             string `json:"type"`
	TypeNumber       int    `json:"typeNumber"`
	Year             int    `json:"year"`
	Description      string `json:"description"`
	ShortDescription string `json:"shortDescription"`
	Slogan           string `json:"slogan"`
	Status           string `json:"status"`
	Rating           struct {
		Kp                 float64 `json:"kp"`
		Imdb               float64 `json:"imdb"`
		Tmdb               float64 `json:"tmdb"`
		FilmCritics        float64 `json:"filmCritics"`
		RussianFilmCritics float64 `json:"russianFilmCritics"`
		Await              float64 `json:"await"`
	} `json:"rating"`
	Votes struct {
		Kp                 int `json:"kp"`
		Imdb               int `json:"imdb"`
		Tmdb               int `json:"tmdb"`
		FilmCritics        int `json:"filmCritics"`
		RussianFilmCritics int `json:"russianFilmCritics"`
		Await              int `json:"await"`
	} `json:"votes"`
	MovieLength int    `json:"movieLength"`
	RatingMpaa  string `json:"ratingMpaa"`
	AgeRating   int    `json:"ageRating"`
	Logo        struct {
		URL string `json:"url"`
	} `json:"logo"`
	Poster struct {
		URL        string `json:"url"`
		PreviewURL string `json:"previewUrl"`
	} `json:"poster"`
	Backdrop struct {
		URL        string `json:"url"`
		PreviewURL string `json:"previewUrl"`
	} `json:"backdrop"`
	Videos struct {
		Trailers []struct {
			URL  string `json:"url"`
			Name string `json:"name"`
			Site string `json:"site"`
			Type string `json:"type"`
			Size int    `json:"size"`
		} `json:"trailers"`
		Teasers []struct {
			URL  string `json:"url"`
			Name string `json:"name"`
			Site string `json:"site"`
			Type string `json:"type"`
			Size int    `json:"size"`
		} `json:"teasers"`
	} `json:"videos"`
	Genres []struct {
		Name string `json:"name"`
	} `json:"genres"`
	Countries []struct {
		Name string `json:"name"`
	} `json:"countries"`
	Persons []struct {
		ID           int    `json:"id"`
		Photo        string `json:"photo"`
		Name         string `json:"name"`
		EnName       string `json:"enName"`
		Description  string `json:"description"`
		Profession   string `json:"profession"`
		EnProfession string `json:"enProfession"`
	} `json:"persons"`
	ReviewInfo struct {
		Count         int    `json:"count"`
		PositiveCount int    `json:"positiveCount"`
		Percentage    string `json:"percentage"`
	} `json:"reviewInfo"`
	SeasonsInfo []struct {
		Number        int `json:"number"`
		EpisodesCount int `json:"episodesCount"`
	} `json:"seasonsInfo"`
	Budget struct {
		Value    int    `json:"value"`
		Currency string `json:"currency"`
	} `json:"budget"`
	Fees struct {
		World struct {
			Value    int    `json:"value"`
			Currency string `json:"currency"`
		} `json:"world"`
		Russia struct {
			Value    int    `json:"value"`
			Currency string `json:"currency"`
		} `json:"russia"`
		Usa struct {
			Value    int    `json:"value"`
			Currency string `json:"currency"`
		} `json:"usa"`
	} `json:"fees"`
	Premiere struct {
		Country string    `json:"country"`
		World   time.Time `json:"world"`
		Russia  time.Time `json:"russia"`
		Digital string    `json:"digital"`
		Cinema  time.Time `json:"cinema"`
		Bluray  string    `json:"bluray"`
		Dvd     string    `json:"dvd"`
	} `json:"premiere"`
	SimilarMovies []struct {
		ID     int `json:"id"`
		Rating struct {
			Kp                 float64 `json:"kp"`
			Imdb               float64 `json:"imdb"`
			Tmdb               float64 `json:"tmdb"`
			FilmCritics        float64 `json:"filmCritics"`
			RussianFilmCritics float64 `json:"russianFilmCritics"`
			Await              float64 `json:"await"`
		} `json:"rating"`
		Year            int    `json:"year"`
		Name            string `json:"name"`
		EnName          string `json:"enName"`
		AlternativeName string `json:"alternativeName"`
		Type            string `json:"type"`
		Poster          struct {
			URL        string `json:"url"`
			PreviewURL string `json:"previewUrl"`
		} `json:"poster"`
	} `json:"similarMovies"`
	SequelsAndPrequels []struct {
		ID     int `json:"id"`
		Rating struct {
			Kp                 float64 `json:"kp"`
			Imdb               float64 `json:"imdb"`
			Tmdb               float64 `json:"tmdb"`
			FilmCritics        float64 `json:"filmCritics"`
			RussianFilmCritics float64 `json:"russianFilmCritics"`
			Await              float64 `json:"await"`
		} `json:"rating"`
		Year            int    `json:"year"`
		Name            string `json:"name"`
		EnName          string `json:"enName"`
		AlternativeName string `json:"alternativeName"`
		Type            string `json:"type"`
		Poster          struct {
			URL        string `json:"url"`
			PreviewURL string `json:"previewUrl"`
		} `json:"poster"`
	} `json:"sequelsAndPrequels"`
	Watchability struct {
		Items []struct {
			Name string `json:"name"`
			Logo struct {
				URL string `json:"url"`
			} `json:"logo"`
			URL string `json:"url"`
		} `json:"items"`
	} `json:"watchability"`
	ReleaseYears []struct {
		Start int `json:"start"`
		End   int `json:"end"`
	} `json:"releaseYears"`
	Top10             int  `json:"top10"`
	Top250            int  `json:"top250"`
	TicketsOnSale     bool `json:"ticketsOnSale"`
	TotalSeriesLength int  `json:"totalSeriesLength"`
	SeriesLength      int  `json:"seriesLength"`
	IsSeries          bool `json:"isSeries"`
	Audience          []struct {
		Count   int    `json:"count"`
		Country string `json:"country"`
	} `json:"audience"`
	Lists    []string `json:"lists"`
	Networks []struct {
		Items []struct {
			Name string `json:"name"`
			Logo struct {
				URL string `json:"url"`
			} `json:"logo"`
		} `json:"items"`
	} `json:"networks"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
	Facts     []struct {
		Value   string `json:"value"`
		Type    string `json:"type"`
		Spoiler bool   `json:"spoiler"`
	} `json:"facts"`
	ImagesInfo struct {
		PostersCount   int `json:"postersCount"`
		BackdropsCount int `json:"backdropsCount"`
		FramesCount    int `json:"framesCount"`
	} `json:"imagesInfo"`
}
