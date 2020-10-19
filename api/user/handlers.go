package user

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

//AddUser Check.
func AddUser(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var userData User
	err := decoder.Decode(&userData)
	if err != nil {
		panic(err)
	}

	prettyJSON, err := json.MarshalIndent(userData, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))

	defer req.Body.Close()

	insertResult, err := UserCollection.InsertOne(context.Background(), userData)
	// insertResult, err := collection.InsertOne(ctx, userData)
	log.Printf(req.URL.Path)
	if err != nil {
		fmt.Printf("%s\n", string(prettyJSON))
		log.Fatal(err)
		res.WriteHeader(500)
		return
	}

	// fmt.Printf("%s\n", string(prettyJSON))
	log.Print("User Insert: ", insertResult.InsertedID)
	response(res, 200, userData)

}

// CreateUser is used to add User Models to DB
func CreateUser(res http.ResponseWriter, request *http.Request) {

	var newUser User
	if err := json.NewDecoder(request.Body).Decode(&newUser); err != nil {
		fmt.Fprintf(res, "Invalid request payload")
		return
	}
	defer request.Body.Close()

	result, err := UserCollection.InsertOne(context.Background(), newUser)
	if err != nil {
		log.Fatal("There is an error")
	}

	log.Print("All DOne Ohk")
	log.Print(result.InsertedID)
	log.Print("User Insert: ", result.InsertedID)
	fmt.Fprintf(res, "DONEEE")
	// response(res, 200, result.InsertedID)
}

//GetUsers Check.
func GetUsers(res http.ResponseWriter, req *http.Request) {

	var allUsers Users
	cursor, err := UserCollection.Find(context.Background(), bson.M{})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var cUser User
		cursor.Decode(&cUser)
		allUsers = append(allUsers, cUser)
	}
	if err := cursor.Err(); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(allUsers)
	// response(res, 200, allUsers)

}

// To have uniform response output
func response(w http.ResponseWriter, status int, results User) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}
