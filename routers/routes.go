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
		userRoutes.GET("/all", controllers.GetUsers)
		userRoutes.POST("/kyc/create", middleware.RequireAuth, controllers.Createkyc)
		userRoutes.GET("/kyc/:id", middleware.RequireAuth, controllers.GetKycByUserID)
	}

	r.GET("/search", controllers.Search)
	r.GET("/search/all", controllers.GlobalSearch)
	r.GET("/search/advanced", controllers.AdvancedSearch)
	r.GET("/search/address", controllers.AddressSearch)
	// r.GET("/search/all/address/{province}/{district}/{municipality}/{ward_number}", controllers.AllAddressSearch)
	r.GET("/search/all/address/:province/:district/:municipality/:ward_number", controllers.AllAddressSearch)

}
