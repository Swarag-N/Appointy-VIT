// +build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	client, error := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	error = client.Connect(ctx)

	//Checking the connection
	error = client.Ping(context.TODO(), nil)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Database connected")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
