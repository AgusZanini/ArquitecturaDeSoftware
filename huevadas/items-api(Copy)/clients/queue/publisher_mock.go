package queue

import (
	"context"
	"items-api/dtos"
)

type PublisherMock struct{}

func (PublisherMock) Publish(ctx context.Context, item dtos.ItemDto) error {
	//TODO implement me
	panic("implement me")
}
