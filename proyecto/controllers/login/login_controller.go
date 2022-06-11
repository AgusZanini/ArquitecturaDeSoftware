package controllers

import (
	"net/http"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/dto"
	"github.com/gin-gonic/gin"
)

/*
func Login(c *gin.Context) {
	var cred dto.Credentials

	if err := c.BindJSON(&cred); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := services.Login(cred)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, token)
}
*/

func login(c *gin.Context) {
	var cred dto.Credentials
	err := c.BindJSON(&cred)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	hash, err := services.login()
}
