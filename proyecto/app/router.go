package app

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var router *gin.engine

func init() {
	router = gin.Default()
}

func StartRoutes() {
	mapUrls()

	log.Info("Starting Server")
	router.Run(":8080")
}
