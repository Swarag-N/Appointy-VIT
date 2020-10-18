package user

import (
	"fmt"
	"log"
	"net/http"
)

// UserAPI is for User.
type UserAPI struct{}

// func (u *UserAPI) ServerHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Hello", r.URL)
// }

func (u *UserAPI) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	switch request.URL.Path {
	case "/api/user/":
		AddUser(response, request)
		break
	case "/api/user/get":
		GetUser(response, request)
		break
	default:
		log.Printf(request.URL.Host)
		log.Printf(request.URL.Path)
		log.Printf(request.URL.RawPath)
		log.Printf(request.URL.RawQuery)
		fmt.Fprintf(response, "Unsupported method '%v' to %v\n", request.Method, request.URL)
		log.Printf("Unsupported method '%v' to %v\n", request.Method, request.URL)
	}
}

//AddUser Check.
func AddUser(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Add User here")
}

//GetUser Check.
func GetUser(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "See User Details")
}
