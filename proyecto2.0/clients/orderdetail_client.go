package clients

import (
	"fmt"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/model"

	"gorm.io/gorm"
)

var Dborderdetail *gorm.DB

func GetOrderDetailsByOrder(id int) model.OrderDetails {
	var orderdetails model.OrderDetails

	if Dborderdetail == nil {
		fmt.Println("nil order detail")
	}

	if result := Dborderdetail.Where("IDorder = ?", id).Find(&orderdetails); result.Error != nil {
		fmt.Println("Couldn't find order details")
	}
	fmt.Println("Orderdetails: ", orderdetails)

	return orderdetails
}

func InsertOrderDetails(orderdetails model.OrderDetails) model.OrderDetails {
	for _, orderdetail := range orderdetails {
		if result := Dborderdetail.Create(&orderdetail); result.Error != nil {
			fmt.Println("Couldn't create orderdetail")
		}
	}

	return orderdetails
}
