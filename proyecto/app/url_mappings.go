package app

import (
	"fmt"

	"fmt"
)

func mapUrls() {
	//mapeo de los clientes

	router.GET("/client/:client_id", client_controller.GetClientById)
	router.GET("/client", client_controller.GetClients)
	router.POST("client", client_controller.InsertClient)

	fmt.Println("Se termino de configurar el mapeo de los datos")
}
