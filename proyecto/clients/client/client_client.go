package clients

import (
	"fmt"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/model"

	//log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Dbclient *gorm.DB

func GetClientById(id int) model.Client {
	var cliente model.Client

	Dbclient.Where("ID_client = ?", id).First(&cliente)
	fmt.Println("Client: ", cliente)

	return cliente
}

func GetClientByUsername(username string) model.Client {
	var cliente model.Client

	Dbclient.Where("Username = ?", username).First(&cliente)
	fmt.Println("Client: ", cliente)

	return cliente
}

func GetClients() model.Clients {
	var clientes model.Clients

	Dbclient.Find(&clientes)
	fmt.Println("clientes: ", clientes)

	return clientes
}

//arreglar
func InsertClient(cliente model.Client) model.Client {
	result := Dbclient.Create(&cliente)

	if result.Error != nil {
		fmt.Println("error")
	}

	fmt.Println("cliente agregado", cliente.ID_client)
	return cliente
}
