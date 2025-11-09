package controllers

import (
	"locpay_api/internal/schemas"
	"locpay_api/internal/services"

	"github.com/gin-gonic/gin"
)

func CreateOperation(c *gin.Context) {

	var data schemas.CreateOperation

	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		c.JSON(400, gin.H{"erro": "campos invalidos"})
		return 
	}

	response, err := services.CreateOperation(&data)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, response)
}

func GetOperationById(c *gin.Context) {

	id := c.Param("id")


	response, err := services.GetOperationById(id) 

	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(200, response)
}

func ConfirmOperation(c *gin.Context) {

	id := c.Param("id")


	response, err := services.ConfirmOperation(id) 

	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(200, response)
}
