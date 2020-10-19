package main

import (
	"appointy/api/meeting"
	"appointy/api/user"

	"log"
	"net/http"
)

func main() {
	// Meeting Rotes are redirected to meeting package
	http.Handle("/api/meetings/", &meeting.Meeting{})
	// User Routes are Redirected to User package
	http.Handle("/api/user/", &user.User{})
	// There is a difference of URL for retriving a Single Meeting
	http.HandleFunc("/api/meeting", meeting.GetMeeting)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
