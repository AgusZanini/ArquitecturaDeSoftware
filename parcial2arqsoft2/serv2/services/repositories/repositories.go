package repositories

import (
	"context"

	dtos "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/dtos"
	e "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/utils/errors"
)

type Repository interface {
	Get(ctx context.Context, id string) (dtos.ItemDto, e.ApiError)
	InsertItem(ctx context.Context, Item dtos.ItemDto) (dtos.ItemDto, e.ApiError)
	Update(ctx context.Context, Item dtos.ItemDto) (dtos.ItemDto, e.ApiError)
	Delete(ctx context.Context, id string) e.ApiError
}
