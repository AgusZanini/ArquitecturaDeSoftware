package app

import (
	"fmt"

	client_controller "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/controllers/client"
	//product_controller "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/controllers/product"

	
)

func mapUrls() {
	//mapeo de los clientes

	router.GET("/client/:client_id", client_controller.GetClientById)
	router.GET("/client", client_controller.GetClients)
	router.POST("client", client_controller.InsertClient)

	fmt.Println("Se termino de configurar el mapeo de los datos")
}
