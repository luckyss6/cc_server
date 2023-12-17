package routes

import (
	"cc_server/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	r.Group("/user")
	r.GET("/get", controller.User.Get)
}
