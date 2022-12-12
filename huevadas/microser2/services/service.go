package services

import (
	dtos "microservicio/dtos"
	e "microservicio/utils/errors"
)

type Service interface{
	
	Get(id string) (dtos.BookDto, e.ApiError)
	 
	Insert(bookDto dtos.BookDto) (dtos.BookDto, e.ApiError)
}

