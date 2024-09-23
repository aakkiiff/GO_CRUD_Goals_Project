package main

import (
	"github.com/gin-gonic/gin"
	"goals.com/db"
	"goals.com/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	db.InitMongoDB()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
