package controllers

import (
	"net/http"

	service "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv1/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Search(c *gin.Context) {
	id := c.Param("id")
	itemsDto, err := service.SearchService.Search(id)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, itemsDto)
		return
	}

	c.JSON(http.StatusOK, itemsDto)
}
