package api

import (
	"fmt"
	"log"
	"net/http"
)

// API act as home route.
type API struct{}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "In api package Fprintf '%v' to %v\n", r.Method, r.URL)
	log.Printf("In api package log printf '%v' to %v\n", r.Method, r.URL)
}

// GetClient Returns a MongoClient when called
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
