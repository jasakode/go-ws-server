package server

import "github.com/gin-gonic/gin"

func Routers(r *gin.RouterGroup) {
	r.GET("/info", GetInfo)
}
