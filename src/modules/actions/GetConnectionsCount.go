package controllers

import (
	"fmt"
	"go-ws-server/src/middleware"
	"syscall"

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

	var info syscall.Sysinfo_t
	err := syscall.Sysinfo(&info)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	count := len(middleware.Sessions)
	c.JSON(200, gin.H{
		"status":  true,
		"message": "Success Get Session Count",
		"data": GetConnectionsCountResponse{
			Count: int64(count),
			ServerInfo: ServerInfo{
				MemoryUnit:  "Mb",
				FreeMemory:  int64(info.Freeram*uint64(info.Unit)) / 1000000,
				TotalMemory: int64(info.Totalram*uint64(info.Unit)) / 1000000,
			},
		},
	})
}
