package services

import (
	"context"
	"net/http"

	dtos "items/dtos"
	e "items/utils/errors"
)

type Service interface {
	Get(ctx context.Context, id string) (dtos.ItemDto, e.ApiError)
	InsertItem(ctx context.Context, Item dtos.ItemDto) (dtos.ItemDto, e.ApiError)
	InsertItems(ctx context.Context, Items dtos.ItemsDto) (dtos.ItemsDto, e.ApiError)
	ValidateRequest(r *http.Request) (*dtos.Claims, error)
}
