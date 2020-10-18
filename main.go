package main

import (
	"appointy/api"
	"appointy/api/user"
	"log"
	"net/http"
)

func main() {
	// http.HandleFunc("/api", Handler)
	http.Handle("/api/meeting", &user.UserAPI{})
	http.Handle("/api/user", &user.UserAPI{})
	http.Handle("/api/", &api.API{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
