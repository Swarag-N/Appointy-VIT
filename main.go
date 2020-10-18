package main

import (
	"appointy/api"
	"appointy/api/user"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func getClient() *mongo.Client {
	uri := "mongodb://localhost:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	return client
}

func main() {
	// http.HandleFunc("/api", Handler)
	// var doc
	// dbClient := getClient()
	// db := dbClient.Database("test")
	// coll := db.Collection("inventory")
	// filter := bson.D{{Key: "color", Value: "Red"}}
	// as, err := coll.FindOne(context.Background(), filter).Decode(&doc)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(as)

	http.Handle("/api/meeting", &user.UserAPI{})
	http.Handle("/api/user/", &user.UserAPI{})
	http.HandleFunc("/api/user/add", user.AddUser)
	http.Handle("/api/", &api.API{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
