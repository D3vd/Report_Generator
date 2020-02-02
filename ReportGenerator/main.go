package main

import (
	"encoding/json"
	"log"
	"strconv"
)

func main() {

	ReportQueuePort := "127.0.0.1:11301"
	ElasticsearchPort := "127.0.0.1:9200"

	var queue Queue
	var es ES
	var s3 S3

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

	if ok := s3.Init(); !ok {
		log.Fatalln("Error while connecting to S3. Check if the credentials are right.")
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

		// Release job if it fails
		if err := json.Unmarshal(jobBody, &reportJob); err != nil {
			log.Println("Error while parsing report job body "+strconv.FormatUint(jobID, 10), err)
			queue.ReleaseJob(jobID)
			continue
		}

		// Convert the UI form time format to time.Time{}
		startTime, ok := ConvertTimeLayoutToISO(reportJob.QueryBody.StartDate)

		// Bury the job since the formating is wrong
		if !ok {
			log.Println("Error while converting  start time layout. Bad Format. Job ID: " + strconv.FormatUint(jobID, 10))
			queue.BuryJob(jobID)
			continue
		}

		endTime, ok := ConvertTimeLayoutToISO(reportJob.QueryBody.EndDate)

		// Bury the job since the formating is wrong
		if !ok {
			log.Println("Error while converting end time layout. Bad Format. Job ID: " + strconv.FormatUint(jobID, 10))
			queue.BuryJob(jobID)
			continue
		}

		// TODO: Make changes here for multiple queries
		// Query ES with Instructions
		hits, ok := es.GetDocumentsWithCarrierAndTimeFrame(
			reportJob.QueryBody.CarrierName,
			startTime,
			endTime,
		)

		// Release jobs if it fails
		if !ok {
			log.Println("Error while querying the database. Job ID: " + strconv.FormatUint(jobID, 10))
			queue.ReleaseJob(jobID)
			continue
		}

		// Parse the hits
		flights, ok := ParseESResultToModel(hits)

		// Release job if it fails
		if !ok {
			log.Println("Error while parsing ES Result. Job ID: " + strconv.FormatUint(jobID, 10))
			queue.ReleaseJob(jobID)
			continue
		}

		// Write the flights models to CSV file
		ok = WriteFlightsToCSV(flights, jobID)

		// Release the job if it fails
		if !ok {
			log.Println("Error while writing the data to CSV. Job ID: " + strconv.FormatUint(jobID, 10))
			queue.ReleaseJob(jobID)
			continue
		}

		_, ok, message := s3.UploadCSVToS3("./output/report" + strconv.FormatUint(jobID, 10) + ".csv")

		if !ok {
			log.Println(message)
			queue.ReleaseJob(jobID)
			continue
		}

		// Delete the job if it was successful
		queue.DeleteJob(jobID)
	}

}
