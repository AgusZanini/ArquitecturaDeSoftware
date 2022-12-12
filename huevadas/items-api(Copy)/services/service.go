package services

import (
	"context"
	"items-api/dtos"
	e "items-api/utils/errors"
)

type Service interface {
	GetItemById(ctx context.Context, id string) (dtos.ItemDto, e.ApiError)
	InsertItem(ctx context.Context, item dtos.ItemDto) (dtos.ItemDto, e.ApiError)
}
