package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jimmy/GinAppp/app/controllers"
)

func RegisterAPIRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/users", controllers.ListUsers)
		api.POST("/users", controllers.CreateUser)
		// Add more API routes here
	}
}
