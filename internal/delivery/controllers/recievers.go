package controllers

import (
	"locpay_api/internal/services"

	"github.com/gin-gonic/gin"
)

func GetRecieverById(c *gin.Context) {

	id := c.Param("id")


	response, err := services.GetRecieverById(id) 

	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(200, response)
}
