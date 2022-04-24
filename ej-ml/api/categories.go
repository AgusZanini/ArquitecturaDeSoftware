package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Categorie struct {
	Id   string `json:"id"`
	Name string `json:"name`
}

type Categories []Categorie

func getcategories(SiteID string) (Categories, error) {

	resp, err := http.Get("https://api.mercadolibre.com/sites/MLA/categories")
	if err != nil {
		return Categories{}, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Categories{}, err
	}
	return parsecategories(bytes)
}

func parsecategories(bytes []byte) (Categories, error) {

	var cats Categories

	err := json.Unmarshal(bytes, &cats)
	if err != nil {
		return Categories{}, err
	}
	return cats, nil
}
