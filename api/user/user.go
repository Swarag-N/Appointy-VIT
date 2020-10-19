package user

import (
	"appointy/database"
	"fmt"
	"log"
	"net/http"
)

// UserCollection in DB
var UserCollection = database.CNX.Database("vitdb").Collection("user")

func (u *User) ServeHTTP(resp http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/api/user/":
		CreateUser(resp, request)
		break
	case "/api/user/get":
		GetUsers(resp, request)
		break
	default:
		fmt.Fprintf(resp, "Unsupported method '%v' to %v\n", request.Method, request.URL)
		log.Printf("Unsupported method '%v' to %v\n", request.Method, request.URL)
	}
}
