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

func (u *UserAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unsupported method '%v' to %v\n", r.Method, r.URL)
	log.Printf("Unsupported method '%v' to %v\n", r.Method, r.URL)
}
