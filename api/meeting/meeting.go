package meeting

import (
	"appointy/database"
	"fmt"
	"log"
	"net/http"
)

// MeetingCollection acts as One connection for meeting MeetingCollection
var MeetingCollection = database.CNX.Database("vitdb").Collection("meeting")

// ServeHTTP function creates http Handler of handling http requests
func (m *Meeting) ServeHTTP(resp http.ResponseWriter, request *http.Request) {
	// Error Handling at route node, for invalid http methods
	switch request.Method {
	case "POST":
		AddMeeting(resp, request)
	case "GET":
		// Routing the request for different Functions
		switch request.URL.Path {
		case "/api/meetings":
			GetMeetingsList(resp, request)
			break
		case "/api/meetings/test":
			TestaddMeeting(resp, request)
			break
		default:
			fmt.Fprintf(resp, " 404 Not Defined'%v' to %v\n", request.Method, request.URL)
			log.Printf("404 Not Defined'%v' to %v\n", request.Method, request.URL)
		}
	default:
		fmt.Fprintf(resp, "Unsupported method '%v' to %v\n", request.Method, request.URL)
		log.Printf("Unsupported method '%v' to %v\n", request.Method, request.URL)
	}
}
