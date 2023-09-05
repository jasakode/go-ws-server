package controllers

import (
	"github.com/gin-gonic/gin"
)

type BodyResponse struct {
	Action  string      `json:"action"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var BlastActions = func(c *gin.Context) {

}
