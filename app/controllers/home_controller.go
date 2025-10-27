package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"app":  "Welcome to GinAppp!",
		"msg":  "Gin + MySQL + Storage + API auto-setup successful 🚀",
	})
}
