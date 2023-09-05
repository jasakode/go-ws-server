package ws

import (
	wscontrollers "go-ws-server/src/modules/ws/ws_controllers"

	"github.com/gin-gonic/gin"
)

func Routers(r *gin.RouterGroup) {
	r.GET("/conections", wscontrollers.CreateSessions)
}
