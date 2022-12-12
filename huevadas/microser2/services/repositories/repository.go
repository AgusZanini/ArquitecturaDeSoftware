package repositories

import (
	"microservicio/dtos"
	"microservicio/utils/errors"
)

type Repository interface {
	Get(id string) (dtos.BookDto, errors.ApiError)
	Insert(dto dtos.BookDto) errors.ApiError // puedo tambien devolver el libro insertado si lo deseo
	Update(dto dtos.BookDto) errors.ApiError
	Delete(id string) errors.ApiError
}


