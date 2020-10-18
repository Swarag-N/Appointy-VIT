package user

import (
	"appointy/database"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// UserAPI is for User.
type UserAPI struct{}

//User DB Model.
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var collection = database.CNX.Database("vitdb").Collection("user")

//Users return List
type Users []User

// var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

// // GetClient is DB Connect
// func GetClient() *mongo.Client {
// 	uri := "mongodb://localhost:27017"
// 	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

// // var collection = &api.GetClient().Database("vitdb").Collection("user")
// // var collection = GetClient().Database("vitdb").Collection("user")

// // func (u *UserAPI) ServerHTTP(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Println("Hello", r.URL)
// }

func (u *UserAPI) ServeHTTP(resp http.ResponseWriter, request *http.Request) {

	switch request.URL.Path {
	case "/api/user/":
		AddUser(resp, request)
		break
	case "/api/user/get":
		GetUsers(resp, request)
		break
	case "/api/user/create":
		createUser(resp, request)
	default:
		log.Printf(request.URL.Host)
		log.Printf(request.URL.Path)
		log.Printf(request.URL.RawPath)
		log.Printf(request.URL.RawQuery)
		fmt.Fprintf(resp, "Unsupported method '%v' to %v\n", request.Method, request.URL)
		log.Printf("Unsupported method '%v' to %v\n", request.Method, request.URL)
	}
}

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

	insertResult, err := collection.InsertOne(context.Background(), userData)
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

func createUser(res http.ResponseWriter, request *http.Request) {

	var newUser User
	if err := json.NewDecoder(request.Body).Decode(&newUser); err != nil {
		fmt.Fprintf(res, "Invalid request payload")
		return
	}
	defer request.Body.Close()

	result, err := collection.InsertOne(context.Background(), newUser)
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

	cursor, err := collection.Find(context.Background(), bson.M{})
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

func response(w http.ResponseWriter, status int, results User) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}
