package controllers

import (
	"net/http"
	"proyectoarqsoft/dtos"
	services "proyectoarqsoft/services/orders"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertOrder(c *gin.Context) {
	var orderrequest dtos.Orderrequestdto

	err := c.BindJSON(&orderrequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	er := services.OrderService.InsertOrder(orderrequest)
	if er != nil {
		c.JSON(http.StatusBadRequest, er.Error())
		return
	}
	c.JSON(http.StatusCreated, orderrequest)
}

func GetOrdersByUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ordersandetailsdto, er := services.OrderService.GetOrdersByUser(id)
	if er != nil {
		c.JSON(http.StatusBadRequest, er.Error())
		return
	}

	c.JSON(http.StatusOK, ordersandetailsdto)
}
