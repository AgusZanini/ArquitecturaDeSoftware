package book

import (
	"context"
	"fmt"

	model "github.com/emikohmann/ucc-arqsoft-2/ej-books/models"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/db"
	"github.com/karlseguin/ccache"
	"github.com/karlseguin/ccache/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetFromCache(id string) model.Book {

	var cache = ccache.New(ccache.Configure())

	book := cache.Get("book:id")


}

func GetById(id string) *Book {
	var book model.Book

	db := db.MongoDb
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return book
	}
	err = db.Collection("books").FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(&book)
	if err != nil {
		fmt.Println(err)
		return book
	}
	return book

}

func Insert(book model.Book) model.Book {
	db := db.MongoDb
	insertBook := book
	insertBook.Id = primitive.NewObjectID()
	_, err := db.Collection("books").InsertOne(context.TODO(), &insertBook)

	if err != nil {
		fmt.Println(err)
		return book
	}
	book.Id = insertBook.Id
	return book
}
