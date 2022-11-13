package repositories

import (
	"context"
	"fmt"

	dtos "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/dtos"
	model "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/models"
	e "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/utils/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepositoryMongoDB struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection string
}

func NewMongoDB(host string, port int, collection string) *RepositoryMongoDB {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port)))
	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}

	names, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}

	fmt.Println("[MongoDB] Initialized connection")
	fmt.Println(fmt.Sprintf("[MongoDB] Available databases: %s", names))

	return &RepositoryMongoDB{
		Client:     client,
		Database:   client.Database("items"),
		Collection: collection,
	}
}

func (repo *RepositoryMongoDB) Get(ctx context.Context, id string) (dtos.ItemDto, e.ApiError) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dtos.ItemDto{}, e.NewBadRequestApiError(fmt.Sprintf("error getting item %s invalid id", id))
	}
	result := repo.Database.Collection(repo.Collection).FindOne(context.TODO(), bson.M{
		"_id": objectID,
	})
	if result.Err() == mongo.ErrNoDocuments {
		return dtos.ItemDto{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
	}
	var item model.Item
	if err := result.Decode(&item); err != nil {
		return dtos.ItemDto{}, e.NewInternalServerApiError(fmt.Sprintf("error getting item %s", id), err)
	}
	return dtos.ItemDto{
		Id:        id,
		Title:     item.Title,
		Price:     item.Price,
		Image:     item.Image,
		Sale_sate: item.Sale_sate,
		Condition: item.Condition,
		Address:   item.Address,
	}, nil
}

func (repo *RepositoryMongoDB) InsertItem(ctx context.Context, item dtos.ItemDto) (dtos.ItemDto, e.ApiError) {
	result, err := repo.Database.Collection(repo.Collection).InsertOne(context.TODO(), model.Item{
		Title:     item.Title,
		Price:     item.Price,
		Image:     item.Image,
		Sale_sate: item.Sale_sate,
		Condition: item.Condition,
		Address:   item.Address,
	})
	if err != nil {
		return dtos.ItemDto{}, e.NewInternalServerApiError(fmt.Sprintf("error inserting item %s", item.Id), err)
	}
	item.Id = fmt.Sprintf(result.InsertedID.(primitive.ObjectID).Hex())
	return item, nil
}

func (repo *RepositoryMongoDB) Update(ctx context.Context, item dtos.ItemDto) (dtos.ItemDto, e.ApiError) {
	_, err := repo.Database.Collection(repo.Collection).UpdateByID(context.TODO(), fmt.Sprintf("%v", item.Id), model.Item{
		Title:     item.Title,
		Price:     item.Price,
		Image:     item.Image,
		Sale_sate: item.Sale_sate,
		Condition: item.Condition,
		Address:   item.Address,
	})
	if err != nil {
		return dtos.ItemDto{}, e.NewInternalServerApiError(fmt.Sprintf("error updating item %s", item.Id), err)
	}
	return item, nil
}

func (repo *RepositoryMongoDB) Delete(ctx context.Context, id string) e.ApiError {
	_, err := repo.Database.Collection(repo.Collection).DeleteOne(context.TODO(), bson.M{"_id": fmt.Sprintf("%s", id)})
	if err != nil {
		return e.NewInternalServerApiError(fmt.Sprintf("error deleting item %s", id), err)
	}
	return nil
}
