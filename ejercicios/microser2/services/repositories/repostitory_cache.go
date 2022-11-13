package repositories

import (
	dtos "github.com/AgusZanini/ArquitecturaDeSoftware/microser2/dtos"
	errors "github.com/AgusZanini/ArquitecturaDeSoftware/microser2/utils/errors"
)

type RespositoryCache struct{}

func (RespositoryCache) Get(id string) (dtos.BookDto, errors.ApiError) {}
