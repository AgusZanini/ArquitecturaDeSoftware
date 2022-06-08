package db

import (
	//"database/sql"

	clientclient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/clients/client"
	productclient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/clients/product"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"fmt"

	"os"

	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "sql10482597"
	DBUser := "sql10482597"
	//DBPass := ""
	DBPass := os.Getenv("MVC_DB_PASS")
	DBHost := "sql10.freemysqlhosting.net"
	// ------------------------

	dsn := DBUser + ":" + DBPass + "@tcp(" + DBHost + ":3306)/" + DBName + "?charset=utf8&parseTime=True"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	clientclient.Dbclient = db
	productclient.Dbproduct = db

}

func StartDbEngine() {
	db.AutoMigrate(&model.Client{})
	db.AutoMigrate(&model.Product{})

	fmt.Println("Finalizando migracion de las tablas de la BD")
}
