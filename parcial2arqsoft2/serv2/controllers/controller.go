package controller

import (
	dtos "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/dtos"
	service "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/services"
	e "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/utils/errors"

	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service service.Service
}

func NewController(service service.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) Get(c *gin.Context) {
	item, apiErr := ctrl.service.Get(c.Request.Context(), c.Param("id"))
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, item)
}

func (ctrl *Controller) Insert(c *gin.Context) {
	var item dtos.ItemDto
	if err := c.BindJSON(&item); err != nil {
		apiErr := e.NewBadRequestApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	item, apiErr := ctrl.service.InsertItem(c.Request.Context(), item)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusCreated, item)
}

/*
func (ctrl *Controller) InsertToQueue(c *gin.Context) {
	var itemsDto dtos.ItemsDto
	err := c.BindJSON(&itemsDto)

	// Error Parsing json param
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	er := ctrl.service.InsertToQueue(itemsDto)

	// Error Queueing
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, itemsDto)
}
*/
