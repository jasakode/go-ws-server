package webhookcontrollers

import (
	"go-ws-server/src/middleware/res"
	"go-ws-server/src/models"

	"github.com/gin-gonic/gin"
)

var GetWebhook = func(c *gin.Context) {
	resp := res.Message(true, "Success Get Webhooks")
	// utils.GetFilterInQuery(models.Webhooks{}, c.Request)
	webhooks, err := models.Webhooks.FindAll(models.Webhooks{})
	if err != nil {
		resp["data"] = nil
		res.Response(c.Writer, resp)
		return
	}
	resp["data"] = webhooks.Webhooks
	resp["count"] = webhooks.Count
	res.Response(c.Writer, resp)
}
