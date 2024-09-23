package main

import (
	"crypto/tls"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goals.com/routes"
)

func initMongoDB() {
	// Setup the mgm default config
	err := mgm.SetDefaultConfig(nil, "goals", options.Client().ApplyURI("mongodb+srv://akif:akif@cluster0.61sku.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetTLSConfig(&tls.Config{InsecureSkipVerify: true}))
	if err != nil {
		panic("could not connect to the database!")
	}
	log.Println("connected to the mongodb database!")
}
func main() {
	server := gin.Default()
	initMongoDB()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
