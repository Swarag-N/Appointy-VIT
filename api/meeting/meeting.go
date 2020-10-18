package meeting

import (
	"fmt"
	"log"
	"net/http"
)

// MeetAPI is for User.
type MeetAPI struct{}

func (u *MeetAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unsupported method '%v' to %v\n", r.Method, r.URL)
	log.Printf("Unsupported method '%v' to %v\n", r.Method, r.URL)
}
