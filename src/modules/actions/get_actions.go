package controllers

import "github.com/gin-gonic/gin"

var GetActions = func(c *gin.Context) {

	c.JSON(200, gin.H{
		"status": true,
	})
}
