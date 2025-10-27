package routes

import "github.com/gin-gonic/gin"

func RegisterWebRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Welcome to your Gin app!",
		})
	})
}
