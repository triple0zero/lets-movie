package botkit

import (
	"fmt"
	"github.com/triple0zero/lets-movie/config"
	"github.com/triple0zero/lets-movie/internal/model"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	validKinopoiskHost string = "kinopoisk.ru"
	apiSchema          string = "https"
	apiHost            string = "api.kinopoisk.dev"
	kinopoiskAPI       string = "v1.4/movie"
	kinopoiskAPIHeader string = "LM_KINOPOISK_HEADER"
	kinopoiskAPIToken  string = "LM_KINOPOISK_TOKEN"
)

func CreateMovieObj(s string) model.Movie {
	log.Printf("[INFO]: На вход получаем ссылку на фильм: %s\n", s)

	url, err := url.ParseRequestURI(s)
	if err != nil {
		log.Println("[ERROR]: Указан неверный формат URL")
	}

	if !strings.Contains(url.Host, validKinopoiskHost) {
		log.Println("[ERROR]: Указаный URL - не кинопоиск")
	}

	var kpIDint int

	kpIDint, _ = strconv.Atoi(strings.Split(url.Path, "/")[2])
	//log.Printf("[INFO]: тип переменой id (%d) кинопоиска: %#v", kpIDint, kpIDint)

	info, err := getRequestToKP(kpIDint)
	if err != nil {
		log.Print("[ERROR]: Не удалось получить информацию от кинопоиска")
	}

	kp := parseKinopoiskResp(info)

	return model.Movie{
		Name:        kp.Name,
		Url:         s,
		Description: kp.Description,
		KpRating:    kp.Rating.Kp,
		ImdbRating:  kp.Rating.Imdb,
	}
}

func getRequestToKP(id int) (*http.Response, error) {

	client := &http.Client{}

	url := fmt.Sprintf("%s://%s/%s/%d", apiSchema, apiHost, kinopoiskAPI, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("[ERROR]: Не удалось создать запрос")
	}
	req.Header.Add(config.GetEnvVariable(kinopoiskAPIHeader), config.GetEnvVariable(kinopoiskAPIToken))

	return client.Do(req)
}

func parseKinopoiskResp(r *http.Response) KinopoiskMovieId {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ERROR]: Не удалось полчить тело ответа", err)
	}
	defer r.Body.Close()

	json, err := ParseJSONFromBytes[KinopoiskMovieId](body)
	if err != nil {
		log.Println("[ERROR]: Can not unmarshal JSON")
		fmt.Println(err)
	}

	log.Println("[Name]:", json.Name)
	log.Println("[Description]:", json.Description)
	log.Println("[Rating kp]:", json.Rating.Kp)
	log.Println("[Rating imdb]:", json.Rating.Imdb)

	return json

}

//type resJSON struct {
//	Id          string `json:"id"`
//	Name        string `json:"name"`
//	Rating      rating `json:"rating"`
//	Description string `json:"description"`
//}
//
//type rating struct {
//	Kp   float64 `json:"kp"`
//	Imdb float64 `json:"imdb"`
//}
