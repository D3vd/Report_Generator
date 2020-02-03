package main

import (
	"log"
	"os"
	"time"

	"github.com/beanstalkd/go-beanstalk"
)

// PushJobToQueue : Pushes Jon JSON to beanstalk queue
func PushJobToQueue(jobJSON []byte) (bool, string, uint64) {

	ReportJobQueue := os.Getenv("REPORT_QUEUE_PORT")

	// Create connection to Beanstalk
	tube, err := beanstalk.Dial("tcp", ReportJobQueue)

	if err != nil {
		errorMessage := "Error while connecting to beanstalk queue."
		log.Println(errorMessage)
		return true, errorMessage, 0
	}

	jobID, err := tube.Put(jobJSON, 0, 0, 20*time.Second)

	if err != nil {
		errorMessage := "Error while pushing to queue."
		log.Println(errorMessage)
		return true, errorMessage, 0
	}

	return false, "Successfully added to queue", jobID
}
