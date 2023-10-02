package userscontrollers

import (
	"encoding/json"
	"go-ws-server/src/middleware/res"
	"go-ws-server/src/models"

	"github.com/gin-gonic/gin"
)

var CreateUser = func(c *gin.Context) {
	var body models.Users
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		resp := res.Message(false, err.Error())
		res.Response(c.Writer, resp)
	}

	resp := res.Message(true, "Successfully created User")
	resp["data"] = ""
	res.Response(c.Writer, resp)
}
