package controllers

import (
	"net/http"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/dto"
	services "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/services"
	"github.com/dgrijalva/jwt-go"
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

/*
func login(c *gin.Context){
	var cred dto.Credentials
	err := c.BindJSON(&cred)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	password, err := services.Login(cred)

	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
	}

	c.JSON(http.StatusOK, password)
}
*/

var jwtkey = []byte("secret_key")

func Login(c *gin.Context) {
	var cred dto.Credentials
	err := c.BindJSON(&cred)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	tokenDto, er := services.Login(cred)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	token, err := jwt.Parse(tokenDto.Token, func(t *jwt.Token) (interface{}, error) { return jwtkey, nil })

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, "firma no valida")
			return
		}
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, "token no valido")
		return
	}

	c.JSON(http.StatusCreated, "se accedio con exito")
}
