package db

import (
	//"os"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/model"
	//clientclient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/clients/client"

	"gorm.io/gorm"

	"fmt"

	"os"

	log "github.com/sirupsen/logrus"
)

var (
	db *gorm.DB
)

func init() {
	// DB Connections Paramters
	DBName := "sql10482597"
	DBUser := "sql10482597"
	//DBPass := ""
	DBPass := os.Getenv("MVC_DB_PASS")
	DBHost := "sql10.freemysqlhosting.net"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	userClient.Db = db
}

func StartDbEngine() {
	db.AutoMigrate(&model.Client{})
	db.AutoMigrate(&model.Product{})

	fmt.Println("Finalizando migracion de las tablas de la BD")
}
