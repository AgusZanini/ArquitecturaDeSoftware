package controllers

import (
	"net/http"
	"proyectoarqsoft/dtos"
	services "proyectoarqsoft/services/users"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userdto, er := services.UserService.GetUserById(id)

	if er != nil {
		c.JSON(http.StatusBadRequest, er.Error())
		return
	}

	c.JSON(http.StatusOK, userdto)
}

func Insert(c *gin.Context) {
	var userdto dtos.Userdto

	err := c.BindJSON(&userdto)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	er := services.UserService.Insert(userdto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	c.JSON(http.StatusCreated, userdto)
}

func Login(c *gin.Context) {
	var cred dtos.Credentials

	err := c.BindJSON(&cred)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responsedto, err := services.UserService.Login(cred)

	c.JSON(http.StatusOK, responsedto)
}
