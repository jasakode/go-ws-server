package users

import (
	userscontrollers "go-ws-server/src/modules/users/users_controllers"

	"github.com/gin-gonic/gin"
)

func Routers(r *gin.RouterGroup) {
	r.GET("/get_user", userscontrollers.GetUsers)
	r.POST("/add_user", userscontrollers.CreateUser)
	r.PATCH("/update_user", userscontrollers.UpdateUser)
	r.DELETE("/delete_user", userscontrollers.DeleteUser)
}
