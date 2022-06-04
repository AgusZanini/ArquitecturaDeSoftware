package clients

import (
	"fmt"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	//usar fmt.print en vez del log.debug
)

var Db *gorm.DB

func GetClientById(id int) model.Client {
	var cliente model.Client

	Db.Where("id = ?", id).First(&cliente)
	log.Debug("Client: ", cliente)

	return cliente
}

func GetClients() model.Clients {
	var clientes model.Clients

	Db.Find(&clientes)
	log.Debug("clientes: ", clientes)

	return clientes
}

//arreglar
func InsertClient(cliente model.Client) model.Client {
	result := Db.Create(&cliente)

	if result.Error != nil {
		fmt.Println("error")
		return
	}
	fmt.Println("cliente agregado: ", cliente.ID_cliente)
	return model.Client
}
