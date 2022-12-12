package clients

import (
	"fmt"
	"proyectoarqsoft/models"

	"gorm.io/gorm"
)

var Db *gorm.DB

type orderdetailClient struct{}

type OrderdetailClientInterface interface {
	InserOrderDetail(orderdetail models.Orderdetail) error
	GetOrderDetailsByOrder(id int) (models.Orderdetails, error)
}

var (
	OrderdetailClient OrderdetailClientInterface
)

func init() {
	OrderdetailClient = &orderdetailClient{}
}

func (o *orderdetailClient) InserOrderDetail(orderdetail models.Orderdetail) error {

	result := Db.Create(&orderdetail)
	if result.Error != nil {
		fmt.Println("no se pudo insertar el detalle de orden")
		return result.Error
	}

	fmt.Println("detalle insertado", orderdetail)
	return nil
}

func (o *orderdetailClient) GetOrderDetailsByOrder(id int) (models.Orderdetails, error) {
	var orderdetails models.Orderdetails
	result := Db.Where("orderid = ?", id).Find(&orderdetails)

	if result.Error != nil {
		fmt.Println("no se pudo obtener los detalles de la orden ", id)
		return models.Orderdetails{}, result.Error
	}

	return orderdetails, nil
}
