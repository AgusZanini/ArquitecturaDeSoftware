package repositories

import (
	"context"
	"items-api/dtos"
	"items-api/utils/errors"
)

type Repository interface {
	GetItemById(ctx context.Context, id string) (dtos.ItemDto, errors.ApiError)
	InsertItem(ctx context.Context, item dtos.ItemDto) (dtos.ItemDto, errors.ApiError)
}
