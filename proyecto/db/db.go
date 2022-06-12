package db

import (
	//"database/sql"

	//clientclient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/clients/client"
	//productclient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/clients/product"
	//orderclient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/clients/order"
	//orderdetailclient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/clients/orderdetail"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"fmt"

	//"os"

	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "basepiola"
	DBUser := "root"
	//DBPass := ""
	DBPass := "root" //os.Getenv("MVC_DB_PASS")
	DBHost := "127.0.0.1"
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
	//clientclient.Dbclient = db
	//productclient.Dbproduct = db
	//orderclient.Dbclient = db
	//orderdetailclient.Dborderdetail = db

}

func StartDbEngine() {
	db.AutoMigrate(&model.Client{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Order{})
	db.AutoMigrate(&model.OrderDetail{})

	fmt.Println("Finalizando migracion de las tablas de la BD")
}
