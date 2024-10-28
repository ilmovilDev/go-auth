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
		auth.POST("/authenticate", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Authenticate user"})
		})
		auth.POST("/forgot-password", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Forgot password"})
		})
		auth.POST("/update-password", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Update password"})
		})
		auth.GET("/confirm-email", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Confirm email"})
		})
		auth.GET("/confirm-token", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Confirm token"})
		})
	}
}
