package controllers

import (
	"net/http"
	"proyectoarqsoft/dtos"
	services "proyectoarqsoft/services/categories"

	"github.com/gin-gonic/gin"
)

func InsertCategory(c *gin.Context) {
	var categorydto dtos.Categorydto
	err := c.BindJSON(&categorydto)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	er := services.CategoryService.InsertCategory(categorydto)

	if er != nil {
		c.JSON(http.StatusBadRequest, er.Error())
		return
	}

	c.JSON(http.StatusCreated, categorydto)
}
