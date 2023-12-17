package routes

import "github.com/gin-gonic/gin"

func InitRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	UserRoute(r)
	return r
}