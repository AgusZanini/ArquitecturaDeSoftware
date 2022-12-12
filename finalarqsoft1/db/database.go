package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	categoryclient "proyectoarqsoft/clients/categories"
	orderdetailsclient "proyectoarqsoft/clients/orderdetails"
	orderclient "proyectoarqsoft/clients/orders"
	productclient "proyectoarqsoft/clients/products"
	userclient "proyectoarqsoft/clients/users"
	"proyectoarqsoft/models"
)

var (
	user   = "root"
	pass   = "root"
	host   = "127.0.0.1"
	port   = 3306
	dbname = "arqsoft1"
)

func Connect() {
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)

	// conexion a la base de datos
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error al conectarse a la base de datos %s\n", dbname)
		log.Fatal(err)
	} else {
		fmt.Printf("Conectado a la base de datos %s\n", dbname)
	}

	userclient.Db = db
	productclient.Db = db
	categoryclient.Db = db
	orderclient.Db = db
	orderdetailsclient.Db = db

	er := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{}, &models.Order{}, models.Orderdetail{})

	if er != nil {
		fmt.Println("Error en la migracion de los datos")
		log.Fatal(er)
	}

	fmt.Println("Migracion de los modelos finalizada")
}
