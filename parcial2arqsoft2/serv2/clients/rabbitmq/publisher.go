package rabbitmq

import (
	"context"

	dtos "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/dtos"
)

type Publisher interface {
	Publish(ctx context.Context, item dtos.ItemDto) error
}
