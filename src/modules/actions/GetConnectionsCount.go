package controllers

import (
	"github.com/gin-gonic/gin"
)

type ServerInfo struct {
	TotalMemory int64
	FreeMemory  int64
	MemoryUnit  string
}

type GetConnectionsCountResponse struct {
	Count      int64      `json:"count"`
	ServerInfo ServerInfo `json:"server_info"`
}

var GetConnectionsCount = func(c *gin.Context) {

}
