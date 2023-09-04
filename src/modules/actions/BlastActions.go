package controllers

import (
	"encoding/json"
	"go-ws-server/src/middleware"
	"go-ws-server/src/ws"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type BodyResponse struct {
	Action  string      `json:"action"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var BlastActions = func(c *gin.Context) {
	var body BodyResponse
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	byteAct, _ := ws.CreateActions[interface{}]("new-order", "order baru").ToByte()
	for _, session := range middleware.Sessions {
		err := session.Socket.WriteMessage(websocket.TextMessage, byteAct)
		if err != nil {
			middleware.DeletedAction(session.Authorization)
		}
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Success Blast Actions",
		"data":    body,
	})
}
