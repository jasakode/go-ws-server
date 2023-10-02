package userscontrollers

import (
	"go-ws-server/src/middleware/res"
	"go-ws-server/src/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := models.Users.FindAll(models.Users{})
	if err != nil {
		resp := res.Message(false, err.Error())
		resp["data"] = nil
		res.Response(c.Writer, resp)
		return
	}

	resp := res.Message(true, "Successfully Get Users")
	resp["count"] = users
	res.Response(c.Writer, resp)
}
