package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/beanstalkd/go-beanstalk"
)

func main() {

	ReportQueuePort := 11301

	// Connect to Beanstalk
	tube, err := beanstalk.Dial("tcp", "127.0.0.1:"+strconv.Itoa(ReportQueuePort))

	if err != nil {
		log.Println("Error while connecting to beanstalk at port " + strconv.Itoa(ReportQueuePort) + ". Make sure the queue is active.")
		return
	}

	for {

		jobID, jobBody, jobReady := GetJobFromQueue(tube)

		if !jobReady {
			continue
		}

		fmt.Println(jobID, string(jobBody))
	}

}
