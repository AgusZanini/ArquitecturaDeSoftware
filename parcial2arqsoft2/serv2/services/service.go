package services

import (
	"context"

	dtos "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/dtos"
	e "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/utils/errors"
)

type Service interface {
	Get(ctx context.Context, id string) (dtos.ItemDto, e.ApiError)
	InsertItem(ctx context.Context, Item dtos.ItemDto) (dtos.ItemDto, e.ApiError)
	//InsertToQueue(Items dtos.ItemsDto) e.ApiError
	//InsertItems(Items dtos.ItemsDto) (dtos.ItemsDto, e.ApiError)
}
