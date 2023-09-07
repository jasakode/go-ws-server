package webhookcontrollers

import (
	"go-ws-server/src/middleware/res"

	"github.com/gin-gonic/gin"
)

var PostWebhook = func(c *gin.Context) {

	resp := res.Message(true, "Success Post Webhook")
	resp["data"] = ""
	res.Response(c.Writer, resp)
}
