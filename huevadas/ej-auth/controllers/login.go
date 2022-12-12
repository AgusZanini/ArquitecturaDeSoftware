package controllers

import (
	"github.com/AgusZanini/ArquitecturaDeSoftware/ej-auth/domain"
	"github.com/AgusZanini/ArquitecturaDeSoftware/ej-auth/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var cred domain.Credentials

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