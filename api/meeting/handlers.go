package meeting

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetMeeting handles to retrive the Individual Meeting
// httpmethod GET
// @param: id: Meeting ObjectID
// returns Meeting
func GetMeeting(res http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		fmt.Fprintf(res, "Unsupported method '%v' to %v\n", req.Method, req.URL)
		return
	}
	// qMeeting acts as an holder to DB retrived Data
	var qMeeting Meeting

	// paramID is the Meeting Used to Retrive the meeting Information
	paramID := req.URL.Query().Get("id")
	// Type Concerstion of (string)paramID to ObjectID
	id, _ := primitive.ObjectIDFromHex(paramID)

	// Mongo Query to retrive the Meeting
	err := MeetingCollection.FindOne(context.Background(), Meeting{ID: id}).Decode(&qMeeting)
	if err != nil {
		log.Fatal("There is an error")
	}

	json.NewEncoder(res).Encode(qMeeting)
	log.Print("Meeting Query: ", id)

}

// AddMeeting to Add meeting to DB by HTTP
func AddMeeting(res http.ResponseWriter, req *http.Request) {
	// newMeeting is used to Cast the data from the req.Body for further usuage
	var newMeeting Meeting

	// Error Handling TO check for all required Parameters
	if err := json.NewDecoder(req.Body).Decode(&newMeeting); err != nil {
		fmt.Fprintf(res, "Invalid request payload")
		return
	}

	// defer the req.Body after used
	defer req.Body.Close()

	result, err := MeetingCollection.InsertOne(context.Background(), newMeeting)
	if err != nil {
		log.Fatal("There is an error")
	}

	log.Print("Meeting Insert: ", result.InsertedID)
	fmt.Fprintf(res, "DONE %s", result.InsertedID)
}

// GetMeetingsList returns the List of Meeting under query Params
func GetMeetingsList(res http.ResponseWriter, req *http.Request) {
	// Logic use QMeeting object to add all the params and make a mongo query
	// mongo quires like like $lte and $gte for Time STamp Comparisions
	// user email id to filter the results

}

// TestaddMeeting Adds Hardcoded Meeting Data.
func TestaddMeeting(res http.ResponseWriter, req *http.Request) {
	// Hardcode Values To add into DB
	result, err := MeetingCollection.InsertOne(context.Background(), bson.D{
		{Key: "Title", Value: "APPPLE"},
		{Key: "EndTime", Value: primitive.Timestamp{T: uint32(time.Now().Unix())}},
		{Key: "CreatedAt", Value: primitive.Timestamp{T: uint32(time.Now().Unix())}},
	})

	if err != nil {
		log.Fatal("There is an error")
	}

	log.Print("Meeting Insert: ", result.InsertedID)
	fmt.Fprintf(res, "DONE %s", result.InsertedID)
}

// To have uniform response output
func response(w http.ResponseWriter, status int, results Meeting) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}
