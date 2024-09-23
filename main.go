package main

import (
	"github.com/gin-gonic/gin"
	"goals.com/db"
	"goals.com/routes"
)

func main() {
	server := gin.Default()
	db.InitMongoDB()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
