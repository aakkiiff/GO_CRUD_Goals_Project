package db

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() {
	// Setup the mgm default config
	err := mgm.SetDefaultConfig(nil, "goals", options.Client().ApplyURI("mongodb+srv://akif:akif@cluster0.61sku.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0&tlsInsecure=true"))
	if err != nil {
		panic("could not connect to the database!")
	}
	log.Println("connected to the mongodb database!")
}
