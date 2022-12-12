package controllers

import (
	"net/http"
	"proyectoarqsoft/dtos"
	services "proyectoarqsoft/services/products"

	"github.com/gin-gonic/gin"
)

func InsertProduct(c *gin.Context) {
	var productdto dtos.Productdto

	if err := c.BindJSON(&productdto); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := services.ProductService.InsertProduct(productdto)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, productdto)
}

func SearchByName(c *gin.Context) {
	query := c.Param("Query")

	productsdto, err := services.ProductService.SearchByName(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if len(productsdto) == 0 {
		c.JSON(http.StatusOK, []dtos.Productdto{})
		return
	}

	c.JSON(http.StatusOK, productsdto)
}

func GetProductsByCategory(c *gin.Context) {
	category := c.Param("name")

	productsdto, err := services.ProductService.GetProductsByCategory(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, productsdto)
}
