package main

import (
	"log"
	"time"

	"github.com/beanstalkd/go-beanstalk"
)

// PushJobToQueue : Pushes Jon JSON to beanstalk queue
func PushJobToQueue(jobJSON []byte) (bool, string, uint64) {

	// Create connection to Beanstalk
	tube, err := beanstalk.Dial("tcp", "127.0.0.1:11300")

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
