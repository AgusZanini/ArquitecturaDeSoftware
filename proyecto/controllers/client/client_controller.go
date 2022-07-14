package controllers

import (
	//"context"
	//"time"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/dto"
	service "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/services"

	//"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/model"
	"github.com/gin-gonic/gin"
	//"github.com/go-playground/validator/v10"

	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)


func GetClientById(c *gin.Context) {
	log.Debug("Id de cliente a cargar: " + c.Param("id")) //preguntar

	id, _ := strconv.Atoi(c.Param("id")) //preguntar
	var clientDto dto.ClientDto

	clientDto, err := service.ClientService.GetClientById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, clientDto)
}

func GetClientByUsername(c *gin.Context) {
	log.Debug("usuario de cliente a cargar: " + c.Param("username")) //preguntar

	username := c.Param("username")
	var clientDto dto.ClientDto

	clientDto, err := service.ClientService.GetClientByUsername(username)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, clientDto)
}

func GetClients(c *gin.Context) {
	var clientsDto dto.ClientsDto
	clientsDto, err := service.ClientService.GetClients()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, clientsDto)
}

func InsertClient(c *gin.Context) {
	var clientDto dto.ClientDto
	err := c.BindJSON(&clientDto) //preguntar

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	clientDto, er := service.ClientService.InsertClient(clientDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, clientDto)
}

/*
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100 *time.Second)
		defer cancel()

		var client model.Client
		err := c.BindJSON(&client)

		if err != nil {
			log.Error(http.StatusBadRequest, err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		validationErr := validate.Struct(client)

		if validationErr != nil {
			log.Error(http.StatusBadRequest, validationErr.Error())
			c.JSON(http.StatusBadRequest, validationErr.Error())
			return
		}

	}
}

*/
