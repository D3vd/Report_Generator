package main

import (
	"fmt"
	"net/http"
	// "github.com/R3l3ntl3ss/Report_generator/Producer"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}

// TestEndpoint ...
func TestEndpoint(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Test")
}
