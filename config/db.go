package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"projects/mongodb-notes-api/controllers"
	"time"
)

func Connect() {

	// Database Config
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // TODO set correct db connection url
	client, err := mongo.NewClient(clientOptions)

	//Context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	//Close connection
	defer cancel()

	// Test db connection
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	// Use database, name = "go-notes-api"
	db := client.Database("go-notes-api")
	controllers.NotesCollection(db)

	return
}
