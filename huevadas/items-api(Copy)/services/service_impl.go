package services

import (
	"context"
	"errors"
	"fmt"
	"io"
	"items-api/clients/queue"
	"items-api/dtos"
	"items-api/services/repositories"
	e "items-api/utils/errors"
	"net/http"
	"os"
)

type ServiceImpl struct {
	localCache repositories.Repository
	distCache  repositories.Repository
	db         repositories.Repository
	queue      queue.Publisher
}

func NewServiceImpl(
	localCache repositories.Repository,
	distCache repositories.Repository,
	db repositories.Repository,
	queue queue.Publisher,
) *ServiceImpl {
	return &ServiceImpl{
		localCache: localCache,
		distCache:  distCache,
		db:         db,
		queue:      queue,
	}
}

func (serv *ServiceImpl) GetItemById(ctx context.Context, id string) (dtos.ItemDto, e.ApiError) {
	var item dtos.ItemDto
	var apiErr e.ApiError
	var source string

	// try to find it in localCache
	item, apiErr = serv.localCache.GetItemById(ctx, id)
	if apiErr != nil {
		if apiErr.Status() != http.StatusNotFound {
			return dtos.ItemDto{}, apiErr
		}
		// try to find it in distCache
		item, apiErr = serv.distCache.GetItemById(ctx, id)
		if apiErr != nil {
			if apiErr.Status() != http.StatusNotFound {
				return dtos.ItemDto{}, apiErr
			}
			// try to find it in db
			item, apiErr = serv.db.GetItemById(ctx, id)
			if apiErr != nil {
				if apiErr.Status() != http.StatusNotFound {
					return dtos.ItemDto{}, apiErr
				} else {
					fmt.Println(fmt.Sprintf("Not found item %s in any datasource", id))
					apiErr = e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
					return dtos.ItemDto{}, apiErr
				}
			} else {
				source = "db"
				defer func() {
					if _, apiErr := serv.distCache.InsertItem(ctx, item); apiErr != nil {
						fmt.Println(fmt.Sprintf("Error trying to save item in distCache %v", apiErr))
					}
					if _, apiErr := serv.localCache.InsertItem(ctx, item); apiErr != nil {
						fmt.Println(fmt.Sprintf("Error trying to save item in localCache %v", apiErr))
					}
				}()
			}
		} else {
			source = "distCache"
			defer func() {
				if _, apiErr := serv.localCache.InsertItem(ctx, item); apiErr != nil {
					fmt.Println(fmt.Sprintf("Error trying to save item in localCache %v", apiErr))
				}
			}()
		}
	} else {
		source = "localCache"
	}

	fmt.Println(fmt.Sprintf("Obtained item from %s!", source))
	return item, nil
}

func (serv *ServiceImpl) InsertItem(ctx context.Context, item dtos.ItemDto) (dtos.ItemDto, e.ApiError) {
	result, apiErr := serv.db.InsertItem(ctx, item)
	if apiErr != nil {
		fmt.Println(fmt.Sprintf("Error inserting item in db: %v", apiErr))
		return dtos.ItemDto{}, apiErr
	}
	fmt.Println(fmt.Sprintf("Inserted item in db: %v", result))

	_, apiErr = serv.distCache.InsertItem(ctx, result)
	if apiErr != nil {
		fmt.Println(fmt.Sprintf("Error inserting item in distCache: %v", apiErr))
		return result, nil
	}
	fmt.Println(fmt.Sprintf("Inserted item in distCache: %v", result))

	_, apiErr = serv.localCache.InsertItem(ctx, result)
	if apiErr != nil {
		fmt.Println(fmt.Sprintf("Error inserting item in localCache: %v", apiErr))
		return result, nil
	}
	fmt.Println(fmt.Sprintf("Inserted item in localCache: %v", result))

	if err := serv.queue.Publish(ctx, result); err != nil {
		return result, e.NewInternalServerApiError(fmt.Sprintf("Error publishing item %s", item.Id), err)
	}
	fmt.Println(fmt.Sprintf("Message sent: %v", result.Id))

	go downloadImage(result.Picture, result.Id)

	return result, nil
}

func downloadImage(url string, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
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
