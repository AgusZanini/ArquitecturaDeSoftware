package controllers

import (
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/dto"
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"

	service "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/services"

	log "github.com/sirupsen/logrus"
)

func GetOrderDetailById(c *gin.Context) {
	log.Debug("id del detalle de orden a cargar" + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var orderdetailDto dto.OrderdetailDto

	orderdetailDto, err := service.Orderdetailservice.GetOrderDetailById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, orderdetailDto)
}

func GetOrderDetails(c *gin.Context) {
	var orderdetailsDto dto.OrderdetailsDto
	orderdetailsDto, err := service.Orderdetailservice.GetOrderDetails()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, orderdetailsDto)
}

func InsertOrderDetail(c *gin.Context) {
	var orderdetailDto dto.OrderdetailDto
	err := c.BindJSON(&orderdetailDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	orderdetailDto, er := service.Orderdetailservice.InsertOrderDetail(orderdetailDto)

	if er != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, orderdetailDto)
}
