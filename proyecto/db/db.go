package db

import (
	//"os"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/model"
	//clientclient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/clients/client"

	"gorm.io/gorm"
	//"gorm.io/driver/mysql"

	"fmt"
)

var (
	db *gorm.DB
)

func StartDbEngine() {
	db.AutoMigrate(&model.Client{})

	fmt.Println("Finalizando migracion de las tablas de la BD")
}
