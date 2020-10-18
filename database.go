package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

func dbconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	client, error := mongo.NewClient(options.Client().ApplyURI("ur_Database_uri"))
	error = client.Connect(ctx)

	//Checking the connection
	error = client.Ping(context.TODO(), nil)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Database connected")

	//Specify your respective collection
	// BooksCollection := client.Database("test").Collection("books")
}
