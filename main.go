package main

import (
	"github.com/BasantaBhusan/go-jwt/contollers"
	"github.com/BasantaBhusan/go-jwt/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", contollers.Signup)
	r.Run()
}
