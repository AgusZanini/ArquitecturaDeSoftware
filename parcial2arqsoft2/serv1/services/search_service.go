package services

import (
	"encoding/json"
	"io/ioutil"

	clients "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv1/clients"
	dtos "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv1/dtos"
)

type searchService struct{}

type searchServiceInterface interface {
	Search(id string) (dtos.ItemsDto, error)
}

var (
	SearchService searchServiceInterface
)

func init() {
	SearchService = &searchService{}
}

func (s *searchService) Search(id string) (dtos.ItemsDto, error) {
	r, err := clients.Search(id)

	if err != nil {
		return dtos.ItemsDto{}, err
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return dtos.ItemsDto{}, err
	}
	rdto, err := parseItems(bytes)

	if err != nil {
		return dtos.ItemsDto{}, err
	}

	return rdto.Docs, nil
}

func parseItems(bytes []byte) (dtos.ResponseDto, error) {
	var rdto dtos.ResponseDto
	if err := json.Unmarshal(bytes, &rdto); err != nil {
		return dtos.ResponseDto{}, err
	}
	return rdto, nil
}
