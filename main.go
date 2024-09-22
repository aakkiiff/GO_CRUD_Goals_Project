package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goals.com/routes"
)

func InitMongoDB() {
	// Setup the mgm default config
	err := mgm.SetDefaultConfig(nil, "note", options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	log.Println("Connected to MongoDB!")
}

func main() {
	server := gin.Default()
	InitMongoDB()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
