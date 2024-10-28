package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	auth := router.Group("/api/auth")
	{
		auth.POST("/register", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Register user"})
		})
	}
}
