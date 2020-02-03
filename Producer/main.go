package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	if reportQueuePort := os.Getenv("REPORT_QUEUE_PORT"); reportQueuePort == "" {
		log.Fatalln("Report Queue Port is not set.")
		return
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/generateReport", GenerateReportJob)
	http.ListenAndServe(":8081", nil)
}
