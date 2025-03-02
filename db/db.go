package db

import (
	"log"
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() {
	// Retrieve environment variables
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	url := os.Getenv("URL")

	// Check if environment variables are set
	if username == "" || password == "" {
		log.Fatal("MongoDB credentials are not set!")
	}

	// Setup the mgm default config
	err := mgm.SetDefaultConfig(nil, "goals", options.Client().ApplyURI(
		"mongodb+srv://"+username+":"+password+"@"+url+"/?retryWrites=true&w=majority&appName=Cluster0te",
	))
	if err != nil {
		panic("could not connect to the database!")
	}
	log.Println("connected to the mongodb database!")
}
