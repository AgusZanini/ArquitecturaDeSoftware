package controllers

import (
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/dto"
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"

	service "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/services"

	log "github.com/sirupsen/logrus"
)

//implementar

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

func GetClients(c *gin.Context) {
	var clientsDto dto.ClientsDto
	clientsDto, err := service.ClientService.GetClients()

	if err != nil {
		c.JSON(err.Status(), err)
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
