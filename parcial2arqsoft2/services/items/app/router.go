package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine, dependencies *Dependencies) {

	router.GET("/items/:id", dependencies.ItemController.Get)
	router.DELETE("/items/user/:userid", dependencies.ItemController.DeleteByUserId)
	// Middleware para validar el token y la solicitud
	//router.Use(dependencies.ItemController.ValidateTokenAndRequest)

	// Products Mapping

	router.POST("/items", dependencies.ItemController.InsertItems)

	fmt.Println("Finishing mappings configurations")
}
