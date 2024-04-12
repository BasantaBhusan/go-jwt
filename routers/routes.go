package routers

import (
	"github.com/BasantaBhusan/go-jwt/controllers"
	"github.com/BasantaBhusan/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/signup", controllers.Signup)
		userRoutes.POST("/login", controllers.Login)
		userRoutes.GET("/validate", middleware.RequireAuth, controllers.Validate)
		userRoutes.GET("/logout", controllers.Logout)
		userRoutes.POST("/kyc/:id", middleware.RequireAuth, controllers.Createkyc)
	}
}
