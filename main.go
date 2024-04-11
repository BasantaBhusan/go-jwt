package main

import (
	"github.com/BasantaBhusan/go-jwt/docs"
	"github.com/BasantaBhusan/go-jwt/initializers"
	"github.com/BasantaBhusan/go-jwt/routers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	docs.SwaggerInfo.BasePath = "/"
	r := gin.Default()
	routers.InitializeRoutes(r)
	url := ginSwagger.URL("http://localhost:3000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url, func(c *ginSwagger.Config) {
		c.DefaultModelsExpandDepth = -1
	}))
	r.Run()
}
