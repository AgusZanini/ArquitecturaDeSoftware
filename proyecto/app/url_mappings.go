package app

import (
	"fmt"

	client_controller "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/controllers/client"
	login_controller "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/controllers/login"
	order_controller "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/controllers/order"
	orderdetail_controller "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/controllers/orderdetail"
	product_controller "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/controllers/product"
)

func mapUrls() {
	//mapeo de los clientes

	router.GET("/client/:client_id", client_controller.GetClientById)
	router.GET("/client/:username", client_controller.GetClientByUsername)
	router.GET("/client", client_controller.GetClients)
	router.POST("client", client_controller.InsertClient)

	//mapeo de los productos

	router.GET("/product/:product_id", product_controller.GetProductById)
	router.GET("/product/:name", product_controller.SearchProducts)
	router.GET("/product", product_controller.GetProducts)
	router.POST("/product", product_controller.InsertProducts)

	//mapeo de las ordenes

	router.GET("/order/:order_id", order_controller.GetOrderById)
	router.GET("/order", order_controller.GetOrders)
	router.POST("/order", order_controller.InsertOrder)

	//mapeo de los detalles de las ordenes

	router.GET("/orderdetail/:orderdetail_id", orderdetail_controller.GetOrderDetailById)
	router.GET("/orderdetail", orderdetail_controller.GetOrderDetails)
	router.POST("/orderdetail", orderdetail_controller.InsertOrderDetail)

	//mapeo del login

	router.POST("/login", login_controller.Login)

	fmt.Println("Se termino de configurar el mapeo de los datos")
}
