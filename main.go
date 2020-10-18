package main

import (
	"appointy/api/meeting"
	"appointy/api/user"

	"log"
	"net/http"
)

//Game is Game.
type Game struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Platform    []string `json:"platform"`
}

// func getClient() *mongo.Client {
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

func main() {
	// http.HandleFunc("/api", Handler)
	// dbClient := getClient()
	// db := dbClient.Database("test")
	// coll := db.Collection("inventory")
	// filter := bson.D{{Key: "color", Value: "Red"}}
	// as, err := coll.FindOne(context.Background(), filter).Decode(&doc)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(as)ds

	// databases := database.CNX.Database("vitdb")
	// coll := databases.Collection("game")
	// result, err := coll.InsertOne(context.Background(), bson.D{
	// 	{Key: "title", Value: "edho okate"},
	// 	{Key: "description", Value: "chav ra rey"},
	// 	{Key: "platform", Value: "naa talkay"},
	// })

	// if err != nil {
	// 	log.Fatal("There is an error")
	// }

	// log.Print("All DOne Ohk")
	// log.Print(result.InsertedID)
	http.Handle("/api/meetings", &meeting.Meeting{})
	http.Handle("/api/user/", &user.UserAPI{})
	// http.HandleFunc("/api/user/add", user.AddUser)
	// http.Handle("/api/", &api.API{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
