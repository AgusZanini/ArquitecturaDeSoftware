package app

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"fmt"
)

var router *gin.engine

func init() {
	router = gin.Default()
}

func StartRoutes() {
	mapUrls()

	fmt.Println("Starting Server")
	router.Run(":8080")
}
