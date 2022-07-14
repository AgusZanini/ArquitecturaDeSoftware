package controllers

import (
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/dto"
	service "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/services"

	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func GetUserById(c *gin.Context) {
	log.Debug("Id de User a buscar: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var UserDto dto.UserDto

	UserDto, err := service.UserService.GetUserById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, UserDto)
}

/*
func GetUserByUsername(c *gin.Context) {
	log.Debug("usuario de User a cargar: " + c.Param("username"))

	username := c.Param("username")
	var UserDto dto.UserDto

	UserDto, err := service.UserService.GetUserByUsername(username)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, UserDto)
}
*/

func GetUsers(c *gin.Context) {
	var UsersDto dto.UsersDto
	UsersDto, err := service.UserService.GetUsers()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, UsersDto)
}

func Login(c *gin.Context) {
	var cred dto.Credentials
	err := c.BindJSON(&cred)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, er := service.UserService.Login(cred)

	if er != nil {
		c.JSON(http.StatusBadRequest, er.Error())
	}

	c.JSON(http.StatusAccepted, token)
}

/*
func InsertUser(c *gin.Context) {
	var UserDto dto.UserDto
	err := c.BindJSON(&UserDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	UserDto, er := service.UserService.InsertUser(UserDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, UserDto)
}
*/
