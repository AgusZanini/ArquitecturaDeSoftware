package app

import (
	"fmt"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/controllers"
)

func mapUrls() {

	//mapeo de los usuarios

	router.GET("/user/:id", controllers.GetUserById)
	router.GET("/user", controllers.GetUsers)
	router.GET("/login", controllers.Login)

	fmt.Println("se completo el mapeo de los datos")
}
