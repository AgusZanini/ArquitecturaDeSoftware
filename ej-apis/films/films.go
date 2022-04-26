package films

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type film struct {
	Id                       string  `json:"id"`
	Title                    string  `json:"title"`
	Original_title           string  `json:"original_title"`
	Original_title_romanised string  `json:"original_title_romanised"`
	Description              string  `json:"description"`
	Director                 string  `json:"director"`
	Producer                 string  `json:"producer"`
	Release_date             string  `json:"release_date"`
	Running_time             float32 `json:"running_time"`
	Rt_score                 float32 `json:"rt_score"`
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

func Getfilms(SiteID string) (Films, error) {
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

func (film film) String() string {
	return fmt.Sprintf(" - %s: %s: %s: %s: %s: %s: %s: %s: %f: %f", film.Id, film.Title, film.Original_title,
		film.Original_title_romanised, film.Description, film.Director,
		film.Producer, film.Release_date, film.Running_time, film.Rt_score)
}
