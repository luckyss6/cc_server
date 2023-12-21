package routes

import (
	"cc_server/controller"

	"github.com/gin-gonic/gin"
)

func DockerRoute(r *gin.Engine) {
	docker := r.Group("/docker")
	{
		docker.GET("/get/images", controller.Docker.GetImages)
		docker.GET("/get/container", controller.Docker.GetContainer)
		docker.POST("/create/container", controller.Docker.CreateContainer)
	}
}
