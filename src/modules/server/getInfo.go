package server

import "github.com/gin-gonic/gin"

var GetInfo = func(c *gin.Context) {

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Success Get Server Info",
		"data":    0,
	})
}
