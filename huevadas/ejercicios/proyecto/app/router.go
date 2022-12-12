package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

func StartRoutes() {
	mapUrls()

	fmt.Println("Starting Server")
	router.Run(":8080")
}
