package Films

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type film struct {
	id                       string  `json:"id"`
	title                    string  `json:"title"`
	original_title           string  `json:"original_title"`
	original_title_romanised string  `json:"original_title_romanised"`
	description              string  `json:"description"`
	director                 string  `json:"director"`
	producer                 string  `json:"producer"`
	release_date             string  `json:"release_date"`
	running_time             float32 `json:"running_time"`
	rt_score                 float32 `json:"rt_score"`
}

type Films []film

func parsefilms(bytes []byte) (Films, error) {
	var films Films
	err := json.Unmarshal(bytes, &films)

	if err != nil {
		return Films{}, err
	}

	return films, nil
}

func getfilms(SiteID string) (Films, error) {
	response, err := http.Get("https://ghibliapi.herokuapp.com/films/58611129-2dbc-4a81-a72f-77ddfc1b1b49")
	if err != nil {
		return Films{}, err
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Films{}, err
	}

	return parsefilms(bytes)
}
