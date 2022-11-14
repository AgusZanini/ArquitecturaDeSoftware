package clients

import (
	"net/http"
)

func Search(id string) (*http.Response, error) {
	url := "http://localhost:8983/solr/items/query?q=id:" + id

	r, err := http.Get(url)

	if err != nil {
		return r, err
	}

	return r, nil
}
