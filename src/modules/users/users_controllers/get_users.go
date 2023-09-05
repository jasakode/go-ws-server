package userscontrollers

import (
	"go-ws-server/src/middleware"
	"go-ws-server/src/middleware/res"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	length := len(middleware.Sessions)
	resp := res.Message(true, "Successfully Get Users")
	resp["count"] = length
	res.Response(c.Writer, resp)
}
