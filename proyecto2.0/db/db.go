package db

import (
	//"database/sql"

	userclient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/clients"
	//orderclient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/clients/order"
	//orderdetailclient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/clients/orderdetail"
	//productclient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/clients/product"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/model"

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
	DBName := "mydb"
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
	userclient.DbUser = db
	//productclient.Dbproduct = db
	//orderclient.Dborder = db
	//orderdetailclient.Dborderdetail = db

}

func StartDbEngine() {
	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{}, &model.OrderDetail{}, &model.Category{})
	fmt.Println("Finalizando migracion de las tablas de la BD")
}

func PopulateDb() {
	user := model.User{Username: "AgusZanini", Firstname: "agustin", Lastname: "zanini", Email: "email", Password: "1234", Phone: 1234}
	db.Select("Username", "Firstname", "Lastname", "Email", "Password", "Phone").Create(&user)
}
