package clients

import (
	"fmt"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/model"

	//log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Dborderdetail *gorm.DB

func GetOrderDetailById(id int) model.OrderDetail {
	var orderdetail model.OrderDetail

	Dborderdetail.Where("id = ?").Find(&orderdetail)
	fmt.Println("orderdetail", orderdetail)

	return orderdetail
}

func GetOrderDetails() model.OrderDetails {
	var orderdetails model.OrderDetails

	Dborderdetail.Find(&orderdetails)
	fmt.Println("detail details", orderdetails)

	return orderdetails
}

func InsertOrderDetail(orderdetail model.OrderDetail) model.OrderDetail {
	result := Dborderdetail.Create(&orderdetail)

	if result.Error != nil {
		fmt.Print("error")
	}

	fmt.Print("detalle de orden agregada", orderdetail.ID_orderdetail)
	return orderdetail
}
