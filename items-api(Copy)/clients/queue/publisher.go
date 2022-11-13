package queue

import (
	"context"
	"items-api/dtos"
)

type Publisher interface {
	Publish(ctx context.Context, item dtos.ItemDto) error
}
