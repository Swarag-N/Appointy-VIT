package api

import (
	"fmt"
	"log"
	"net/http"
)

// API act as home route.
type API struct{}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "In api package Fprintf '%v' to %v\n", r.Method, r.URL)
	log.Printf("In api package log printf '%v' to %v\n", r.Method, r.URL)
}
