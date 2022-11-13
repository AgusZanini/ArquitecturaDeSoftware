package clients

import (
	"fmt"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/model"

	"gorm.io/gorm"
)

var Dborder *gorm.DB

func GetOrdersByUser(id int) model.Orders {
	var orders model.Orders

	if result := Dborder.Where("IDuser = ?", id).Find(&orders); result.Error != nil {
		fmt.Println("Couln't find orders for given user")
	}
	fmt.Println("Orders: ", orders)

	return orders
}

func GetOrders() model.Orders {
	var orders model.Orders

	if result := Dborder.Find(&orders); result.Error != nil {
		fmt.Println("Couln't find orders")
	}
	fmt.Println("Orders: ", orders)

	return orders
}

func InsertOrder(order model.Order) model.Order {

	if result := Dborder.Create(&order); result.Error != nil {
		fmt.Println("couldn't create order")
	}
	fmt.Println("Order created", order)

	return order
}

func UpdateTotalPrice(id int, totalprice float32) {
	if result := Dborder.Model(&model.Order{}).Where("id = ?", id).Update("Totalprice", totalprice); result.Error != nil {
		fmt.Println("couldn't update order")
	}
}
