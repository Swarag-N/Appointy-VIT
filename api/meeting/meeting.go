package meeting

import (
	"appointy/database"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO Change RSVP

//  00 Not Answered
//  01 MayBe
//  10 No
//  11 Yes

//  Done

var collection = database.CNX.Database("vitdb").Collection("meeting")

// Participant Object in Meeting.
type Participant struct {
	UID   primitive.ObjectID `json:"uid" bson:"uid"`
	Email string             `json:"email" bson:"email"`
	rsvp  string             `json:"rsvp" bson:"rsvp`
}

// Meeting is for User.
type Meeting struct {
	ID           primitive.ObjectID  `json:"_id" bson:"_id"`
	Title        string              `json:"title" bson:"title"`
	StartTime    primitive.Timestamp `json:"startTime" bson:"startTime"`
	EndTime      primitive.Timestamp `json:"endTime" bson:"endTime"`
	CreatedAt    primitive.Timestamp `json:"created_at" bson:"created_at`
	Participants []Participant       `json:"Participants" bson:"Participants"`
}

func (u *Meeting) ServeHTTP(resp http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		addMeeting(resp, request)
	case "GET":
		switch request.URL.Path {
		case "/api/meetings":
			testaddMeeting(resp,request)
			break
		default:
			// log.Printf(request.URL.Host)
			// log.Printf(request.URL.Path)
			// log.Printf(request.URL.RawPath)
			// log.Printf(request.URL.RawQuery)
			fmt.Fprintf(resp, " 404 Not Defined'%v' to %v\n", request.Method, request.URL)
			log.Printf("404 Not Defined'%v' to %v\n", request.Method, request.URL)
		}
	default:
		fmt.Fprintf(resp, "Unsupported method '%v' to %v\n", request.Method, request.URL)
		log.Printf("Unsupported method '%v' to %v\n", request.Method, request.URL)
	}
}

func addMeeting(res http.ResponseWriter, req *http.Request) {
	var newMeeting Meeting
	if err := json.NewDecoder(req.Body).Decode(&newMeeting); err != nil {
		fmt.Fprintf(res, "Invalid request payload")
		return
	}
	defer req.Body.Close()

	result, err := collection.InsertOne(context.Background(), newMeeting)
	if err != nil {
		log.Fatal("There is an error")
	}

	log.Print("Meeting Insert: ", result.InsertedID)
	fmt.Fprintf(res, "DONEEE %s", result.InsertedID)
}

func getMeetingsList(res http.ResponseWriter, req *http.Request) {

}

// GetMeeting Individual Meeting
func GetMeeting(res http.ResponseWriter, req *http.Request) {

}

func testaddMeeting(res http.ResponseWriter, req *http.Request) {

	result, err := collection.InsertOne(context.Background(), bson.D{
		{Key:"ID",Value:""}
		{Key:"Title",Value:""}
		{Key:"StartTime",Value:""}
		{Key:"Participants",Value:""}
		{Key:"EndTime",Value:""}
		{Key:"CreatedAt",Value:""}
	})
	if err != nil {
		log.Fatal("There is an error")
	}

	log.Print("Meeting Insert: ", result.InsertedID)
	fmt.Fprintf(res, "DONEEE %s", result.InsertedID)
}
