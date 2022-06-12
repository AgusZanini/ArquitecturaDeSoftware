package product

import (
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/dto"
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"

	service "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/services"

	log "github.com/sirupsen/logrus"
)

func GetProductById(c *gin.Context) {
	log.Debug("Id de producto a buscar: " + c.Param("id")) //preguntar

	id, _ := strconv.Atoi(c.Param("id")) //preguntar

	var productDto dto.ProductDto
	productDto, err := service.Productservice.GetProductById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productDto)
}

func GetProducts(c *gin.Context) {
	var productsDto dto.ProductsDto
	productsDto, err := service.Productservice.GetProducts()

	if err != nil {
		log.Error(err.Error())
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}

func SearchProducts(c *gin.Context) {
	log.Debug("nombre de producto a buscar" + c.Param("name"))
	name := c.Param("name")

	var productsDto dto.ProductsDto
	productsDto, err := service.Productservice.SearchProducts(name)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, productsDto)
}

func InsertProducts(c *gin.Context) {
	var productDto dto.ProductDto
	err := c.BindJSON(&productDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	productDto, er := service.Productservice.InsertProducts(productDto)

	if er != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, productDto)
}
