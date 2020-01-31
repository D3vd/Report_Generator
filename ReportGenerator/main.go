package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func main() {

	ReportQueuePort := "127.0.0.1:11301"

	var queue Queue

	if ok := queue.Init(ReportQueuePort); !ok {
		log.Println("Error while connecting to beanstalk at port " + ReportQueuePort + ". Make sure the queue is active.")
		return
	}

	defer queue.CloseQueue()

	for {

		jobID, jobBody, jobReady := queue.GetJobFromQueue()

		if !jobReady {
			continue
		}

		var reportJob ReportJob

		if err := json.Unmarshal(jobBody, &reportJob); err != nil {
			log.Println("Error while parsing report job body "+strconv.FormatUint(jobID, 10), err)
			queue.ReleaseJob(jobID)
			continue
		}

		fmt.Println(reportJob)
	}

}
