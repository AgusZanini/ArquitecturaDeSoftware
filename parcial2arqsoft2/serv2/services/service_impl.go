package services

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	//log "github.com/sirupsen/logrus"

	"github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/clients/rabbitmq"
	dtos "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/dtos"
	"github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/services/repositories"
	e "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/utils/errors"
)

type ServiceImpl struct {
	localCache repositories.Repository
	db         repositories.Repository
	queue      rabbitmq.Publisher
	//queue      clients.QueueClient
}

func NewServiceImpl(
	localCache repositories.Repository,
	db repositories.Repository,
	queue rabbitmq.Publisher,
	//queue clients.QueueClient,
) *ServiceImpl {
	return &ServiceImpl{
		localCache: localCache,
		db:         db,
		queue:      queue,
	}
}

func (serv *ServiceImpl) Get(ctx context.Context, id string) (dtos.ItemDto, e.ApiError) {
	var item dtos.ItemDto
	var apiErr e.ApiError
	var source string

	// try to find it in localCache
	item, apiErr = serv.localCache.Get(ctx, id)
	if apiErr != nil {
		if apiErr.Status() != http.StatusNotFound {
			return dtos.ItemDto{}, apiErr
		}
		// try to find it in db
		item, apiErr = serv.db.Get(ctx, id)
		if apiErr != nil {
			if apiErr.Status() != http.StatusNotFound {
				return dtos.ItemDto{}, apiErr
			} else {
				fmt.Printf("Not found item %s in any datasource", id)
				apiErr = e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
				return dtos.ItemDto{}, apiErr
			}
		} else {
			source = "db"
			defer func() {
				if _, apiErr := serv.localCache.InsertItem(ctx, item); apiErr != nil {
					fmt.Printf("Error trying to save item in localCache %v", apiErr)
				}
			}()
		}
	} else {
		source = "localCache"
	}

	fmt.Printf("Obtained item from %s!", source)
	return item, nil
}

func downloadImage(url string, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return e.NewNotFoundApiError("image not found")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func (serv *ServiceImpl) InsertItem(ctx context.Context, item dtos.ItemDto) (dtos.ItemDto, e.ApiError) {
	result, apiErr := serv.db.InsertItem(ctx, item)
	if apiErr != nil {
		fmt.Println(fmt.Sprintf("Error inserting item in db: %v", apiErr))
		return dtos.ItemDto{}, apiErr
	}
	fmt.Println(fmt.Sprintf("Inserted item in db: %v", result))

	if err := serv.queue.Publish(ctx, result); err != nil {
		return result, e.NewInternalServerApiError(fmt.Sprintf("Error publishing item %s", item.Id), err)
	}
	fmt.Println(fmt.Sprintf("Message sent: %v", result.Id))

	go downloadImage(result.Image, result.Id)

	return result, nil
}

/*
func (serv *ServiceImpl) InsertToQueue(itemsDto dtos.ItemsDto) e.ApiError {
	for i := range itemsDto {
		var item dtos.ItemDto

		item = itemsDto[i]
		go func() {
			item, err := serv.db.InsertItem(item)
			if err != nil {
				log.Debug(err)
			}
			err = serv.queue.SendMessage("solr", item.Id)
			log.Debug(err)
		}()

	}
	return nil
}
*/
