package modules

import (
	"go-ws-server/src/modules/server"
	"go-ws-server/src/modules/users"
	"go-ws-server/src/modules/webhook"
	"go-ws-server/src/modules/ws"

	"github.com/gin-gonic/gin"
)

var Routers = func(r *gin.RouterGroup) {
	server.Routers(r.Group("/server"))
	webhook.Routers(r.Group("/webhook"))
	ws.Routers(r.Group("/ws"))
	users.Routers(r.Group("/user"))
}
