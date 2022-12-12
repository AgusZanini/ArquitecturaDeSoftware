package repositories

import (
	dtos "microservicio/dtos"
	errors "microservicio/utils/errors"
)

type RespositoryCache struct{}

func (RespositoryCache) Get(id string) (dtos.BookDto, errors.ApiError) {}
