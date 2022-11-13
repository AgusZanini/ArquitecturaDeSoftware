package main

import (
	"github/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.POST("/login", controllers.login)
	engine.Run(":8080")
}
