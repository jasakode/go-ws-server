package modules

import (
	"go-ws-server/src/modules/server"
	"go-ws-server/src/modules/webhook"

	"github.com/gin-gonic/gin"
)

var Routers = func(r *gin.RouterGroup) {
	server.Routers(r.Group("/server"))
	webhook.Routers(r.Group("/webhook"))
}
