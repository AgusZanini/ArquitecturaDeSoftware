package app

import (
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/controllers/client_controller"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	//mapeo de los clientes

	router.GET("/client/:client_id", client_controller.GetClientById)
	router.GET("/client", client_controller.GetClients)
	router.POST("client", client_controller.InsertClient)

	log.Info("Se termino de configurar el mapeo de los datos")
}
