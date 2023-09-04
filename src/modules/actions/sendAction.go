package controllers

import (
	"go-ws-server/src/middleware"
	"go-ws-server/src/ws"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var SendAction = func(c *gin.Context) {
	byteAct, _ := ws.CreateActions[interface{}]("new-order", "order baru").ToByte()

	for conn := range middleware.Connections {
		err := conn.WriteMessage(websocket.TextMessage, byteAct)
		if err != nil {
			log.Println(err)
			conn.Close()
			delete(middleware.Connections, conn)
		}
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Success Send Action",
		"data":    nil,
	})
}
