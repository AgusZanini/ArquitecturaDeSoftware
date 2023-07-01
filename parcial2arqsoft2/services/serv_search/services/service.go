package services

import (
	"search/dtos"
)

type Service interface {
	Search(query string) (dtos.ItemsDto, error)
	SearchByUserId(id int) (dtos.ItemsDto, error)
}
