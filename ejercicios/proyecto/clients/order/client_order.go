package clients

import (
	"fmt"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/model"

	//log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Dborder *gorm.DB

func GetOrderById(id int) model.Order {
	var order model.Order

	Dborder.Where("ID_order = ?", id).First(&order)
	fmt.Println("order", order)

	return order
}

func GetOrders() model.Orders {
	var orders model.Orders

	Dborder.Find(&orders)
	fmt.Println("orders", orders)

	return orders
}

func InsertOrder(order model.Order) model.Order {
	result := Dborder.Create(&order)

	if result.Error != nil {
		fmt.Print("error")
	}

	fmt.Print("orden agregada", order.ID_order)
	return order
}
