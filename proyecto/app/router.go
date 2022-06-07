package app

import (
	"github.com/gin-gonic/gin"
	"fmt"
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
