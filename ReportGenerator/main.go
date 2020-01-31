package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func main() {

	ReportQueuePort := "127.0.0.1:11301"
	ElasticsearchPort := "127.0.0.1:9200"

	var queue Queue
	var es ES

	// Connect to Beanstalk
	if ok := queue.Init(ReportQueuePort); !ok {
		log.Fatalln("Error while connecting to beanstalk at port " + ReportQueuePort + ". Make sure the queue is active.")
		return
	}
	defer queue.CloseQueue()

	// Connect to Elasticsearch
	if ok := es.Init(ElasticsearchPort); !ok {
		log.Fatalln("Error while connecting to elasticsearch at port " + ElasticsearchPort + ". Make sure that elasticsearch has been started.")
		return
	}

	// Infinite for loop to go through all the jobs in the queue
	for {
		// Get job from Queue
		jobID, jobBody, jobReady := queue.GetJobFromQueue()

		// If there is no job in the queue then try again
		if !jobReady {
			continue
		}

		// Unmarshal report job
		var reportJob ReportJob

		if err := json.Unmarshal(jobBody, &reportJob); err != nil {
			log.Println("Error while parsing report job body "+strconv.FormatUint(jobID, 10), err)
			queue.ReleaseJob(jobID)
			continue
		}

		fmt.Println(reportJob)
	}

}
