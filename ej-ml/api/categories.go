package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Categories []category

func Getcategories(SiteID string) (Categories, error) {

	resp, err := http.Get("https://api.mercadolibre.com/sites/MLA/categories")
	if err != nil {
		return Categories{}, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Categories{}, err
	}
	return Parsecategories(bytes)
}

func Parsecategories(bytes []byte) (Categories, error) {

	var cats Categories

	err := json.Unmarshal(bytes, &cats)
	if err != nil {
		return Categories{}, err
	}
	return cats, nil
}

func (category category) String() string {
	return fmt.Sprintf(" - %s: %s", category.Id, category.Name)
}
