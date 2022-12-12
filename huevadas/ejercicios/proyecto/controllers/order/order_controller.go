package controllers

import (
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/dto"
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"

	service "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/services"

	log "github.com/sirupsen/logrus"
)

func GetOrderById(c *gin.Context) {
	log.Debug("ide de la orden a cargar" + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))

	var orderDto dto.OrderDto
	orderDto, err := service.OrderService.GetOrderById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, orderDto)
}

func GetOrders(c *gin.Context) {
	var ordersDto dto.OrdersDto
	ordersDto, err := service.OrderService.GetOrders()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, ordersDto)
}

func InsertOrder(c *gin.Context) {
	var orderDto dto.OrderDto
	err := c.BindJSON(&orderDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	orderDto, er := service.OrderService.InsertOrder(orderDto)

	if er != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, orderDto)
}
