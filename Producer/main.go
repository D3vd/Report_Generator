package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", TestEndpoint)
	http.ListenAndServe(":8080", nil)
}

// TestEndpoint ...
func TestEndpoint(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Test")
}
