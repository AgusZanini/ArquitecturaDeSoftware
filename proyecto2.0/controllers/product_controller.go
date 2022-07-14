package controllers

import (
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/dto"
	service "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/services"

	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func GetProductById(c *gin.Context) {
	log.Debug("id de product a buscar: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))

	var productdto dto.ProductDto
	productdto, err := service.ProductService.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, productdto)
}

func GetProducts(c *gin.Context) {
	var productsdto dto.ProductsDto
	productsdto, err := service.ProductService.GetProducts()

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, productsdto)
}

func GetProductByCategory(c *gin.Context) {
	log.Debug("id de categoria a buscar" + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))

	var productsdto dto.ProductsDto
	productsdto, err := service.ProductService.GetProductByCategory(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, productsdto)
}

func SearchProduct(c *gin.Context) {
	log.Debug("buscar producto: " + c.Param("name"))
	name := c.Param("name")

	var productsdto dto.ProductsDto
	productsdto, err := service.ProductService.SearchProduct(name)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, productsdto)
}
