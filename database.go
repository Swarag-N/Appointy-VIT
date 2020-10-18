package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type meetingType struct {
	name      string
	startTime time.Time
	endTime   time.Time
}

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

// db := client.Database("test")
// 	coll := db.Collection("inventory")
// 	result, err := coll.InsertOne(
// 		context.Background(),
// 		bson.D{
// 			{"item", "canvas"},
// 			{"qty", 100},
// 			{"tags", bson.A{"cotton"}},
// 			{"size", bson.D{
// 				{"h", 28},
// 				{"w", 35.5},
// 				{"uom", "cm"},
// 			}},
// 		})

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(result)

// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
// if err != nil {
// 	log.Fatal(err)
// 	return nil
// }
// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// defer cancel()
// err = client.Connect(ctx)
// err = client.Ping(context.TODO(), nil)
// fmt.Println("Connected to Mongo")
