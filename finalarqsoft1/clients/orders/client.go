package clients

import (
	"fmt"
	"proyectoarqsoft/models"

	"gorm.io/gorm"
)

var Db *gorm.DB

type orderClient struct{}

type OrderClientInterface interface {
	InsertOrder(order models.Order) error
	Getordersbyuser(id int) (models.Orders, error)
	GetLastOrder() (models.Order, error)
}

var OrderClient OrderClientInterface

func init() {
	OrderClient = &orderClient{}
}

func (o *orderClient) InsertOrder(order models.Order) error {
	result := Db.Create(&order)
	if result.Error != nil {
		fmt.Println("error al insertar la orden")
		return result.Error
	}
	return nil
}

func (o *orderClient) Getordersbyuser(id int) (models.Orders, error) {
	var orders models.Orders

	result := Db.Where("userid = ?", id).Find(&orders)
	if result.Error != nil {
		return models.Orders{}, result.Error
	}

	return orders, nil
}

func (o *orderClient) GetLastOrder() (models.Order, error) {
	var order models.Order

	result := Db.Last(&order)
	if result.Error != nil {
		return models.Order{}, result.Error
	}

	return order, nil
}
