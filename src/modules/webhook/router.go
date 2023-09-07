package webhook

import (
	webhookcontrollers "go-ws-server/src/modules/webhook/webhook_controllers"

	"github.com/gin-gonic/gin"
)

func Routers(r *gin.RouterGroup) {
	r.GET("/get_webhook", webhookcontrollers.GetWebhook)
	r.POST("/post_webhook/*any", webhookcontrollers.PostWebhook)
}
