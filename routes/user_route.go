package routes

import (
	"cc_server/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/login", controller.User.Login)
	}
}
