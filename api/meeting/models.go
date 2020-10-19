package meeting

import "go.mongodb.org/mongo-driver/bson/primitive"

// Participant is stored as an array in Meeting.
// This makes it easy to query using email id
// or else there is need to populate results and then query
type Participant struct {
	UID   primitive.ObjectID `json:"uid" bson:"uid"`
	Email string             `json:"email" bson:"email"`
	// TODO
	// Change RSVP to Binary
	//	00 Not Answered
	//  01 MayBe
	//  10 No
	//  11 Yes
	rsvp string `json:"rsvp" bson:"rsvp"`
}

// A Meeting represents the MongoDB Model of Meeting
type Meeting struct {
	ID           primitive.ObjectID  `json:"_id" bson:"_id"`
	Title        string              `json:"title" bson:"title"`
	StartTime    primitive.Timestamp `json:"startTime" bson:"startTime"`
	EndTime      primitive.Timestamp `json:"endTime" bson:"endTime"`
	CreatedAt    primitive.Timestamp `json:"created_at" bson:"created_at`
	Participants []Participant       `json:"Participants" bson:"Participants"`
}

//MQuery represents the various fields
// considered while Querying meetings
type MQuery struct {
	ID        primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	StartTime primitive.Timestamp `json:"startTime,omitempty,string" bson:"startTime,omitempty"`
	EndTime   primitive.Timestamp `json:"endTime,omitempty,string" bson:"endTime,omitempty"`
	Email     string              `json:"email" bson:"email"`
}
