package database

import "go.mongodb.org/mongo-driver/mongo"

// Database contains mongo.Database
type Database struct {
	database *mongo.Database
}

// Collection returns database
func (d *Database) Collection(collection string) *Collection {
	return &Collection{collection: d.database.Collection(collection)}
}

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"go.mongodb.org/mongo-driver/mongo/readpref"
// )

// type database struct{}

// // GetClient Returns a MongoCLient when called
// func GetClient() *mongo.Client {
// 	uri := "mongodb://localhost:27017"
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer func() {
// 		if err = client.Disconnect(ctx); err != nil {
// 			panic(err)
// 		}
// 	}()
// 	// Ping the primary
// 	if err := client.Ping(ctx, readpref.Primary()); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Successfully connected and pinged.")
// 	return client
// }
