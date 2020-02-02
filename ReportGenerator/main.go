package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

func main() {

	if ok, missing := EnsureENVSet(); !ok {
		log.Fatalln(missing + " is not set. Set all ENV to continue")
		return
	}

	ReportQueuePort := os.Getenv("REPORT_QUEUE_PORT")
	NotifierQueuePort := os.Getenv("NOTIFIER_QUEUE_PORT")
	ElasticsearchPort := os.Getenv("ES_PORT")

	var reportQ Queue
	var notifierQ Queue
	var es ES
	var s3 S3

	// Connect to Report Queue
	if ok := reportQ.Init(ReportQueuePort); !ok {
		log.Fatalln("Error while connecting to Report Queue at port " + ReportQueuePort + ". Make sure the queue is active.")
		return
	}
	defer reportQ.CloseQueue()

	// Connect to Notifier Queue
	if ok := notifierQ.Init(NotifierQueuePort); !ok {
		log.Fatalln("Error while connecting to Notifier Queue at port " + NotifierQueuePort + ". Make sure the queue is active.")
		return
	}
	defer notifierQ.CloseQueue()

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
		jobID, jobBody, jobReady := reportQ.GetJobFromQueue()

		// If there is no job in the queue then try again
		if !jobReady {
			continue
		}

		// Unmarshal report job
		var reportJob ReportJob

		// Release job if it fails
		if err := json.Unmarshal(jobBody, &reportJob); err != nil {
			log.Println("Error while parsing report job body "+strconv.FormatUint(jobID, 10), err)
			reportQ.BuryJob(jobID)
			continue
		}

		// Convert the UI form time format to time.Time{}
		startTime, ok := ConvertTimeLayoutToISO(reportJob.QueryBody.StartDate)

		// Bury the job since the formating is wrong
		if !ok {
			log.Println("Error while converting  start time layout. Bad Format. Job ID: " + strconv.FormatUint(jobID, 10))
			reportQ.BuryJob(jobID)
			continue
		}

		endTime, ok := ConvertTimeLayoutToISO(reportJob.QueryBody.EndDate)

		// Bury the job since the formating is wrong
		if !ok {
			log.Println("Error while converting end time layout. Bad Format. Job ID: " + strconv.FormatUint(jobID, 10))
			reportQ.BuryJob(jobID)
			continue
		}

		// TODO: Make changes here for multiple queries
		// Query ES with Instructions
		hits, totalHits, ok := es.GetDocumentsWithCarrierAndTimeFrame(
			reportJob.QueryBody.CarrierName,
			startTime,
			endTime,
		)

		// Release jobs if it fails
		if !ok {
			log.Println("Error while querying the database. Job ID: " + strconv.FormatUint(jobID, 10))
			reportQ.ReleaseJob(jobID)
			continue
		}

		// Parse the hits
		flights, ok := ParseESResultToModel(hits)

		// Release job if it fails
		if !ok {
			log.Println("Error while parsing ES Result. Job ID: " + strconv.FormatUint(jobID, 10))
			reportQ.ReleaseJob(jobID)
			continue
		}

		// Write the flights models to CSV file
		ok = WriteFlightsToCSV(flights, jobID)

		// Release the job if it fails
		if !ok {
			log.Println("Error while writing the data to CSV. Job ID: " + strconv.FormatUint(jobID, 10))
			reportQ.ReleaseJob(jobID)
			continue
		}

		// Upload the CSV file to S3
		fileURL, ok, message := s3.UploadCSVToS3(jobID, reportJob.UserInfo.Name)

		// If uploads fails then delete the CSV file and release the JOb
		if !ok {
			log.Println(message)
			DeleteCSVFile(jobID)
			reportQ.ReleaseJob(jobID)
			continue
		}

		// Delete the report csv after it's been successfully uploaded
		if ok := DeleteCSVFile(jobID); !ok {
			log.Println("Unable to delete report for Job ID : " + strconv.FormatUint(jobID, 10))
		}

		// Create Notifier Job
		notifierJob := NotifierJob{
			User{
				Name:  reportJob.UserInfo.Name,
				Email: reportJob.UserInfo.Email,
			},
			Search{
				TotalHits: totalHits,
			},
			Result{
				URL: fileURL,
			},
		}

		// Marshal Notifier Job
		notifierJobJSON, err := json.Marshal(notifierJob)

		// If the Marshalling fails then and release the job
		if err != nil {
			log.Println("Error Marshalling Notifier Job")
			reportQ.ReleaseJob(jobID)
			continue
		}

		// Push the Notifier Job to the Notifier Queue
		if ok := notifierQ.PutJob(notifierJobJSON); !ok {
			log.Println("Error While pushing job to the Notifier Queue Job ID: " + strconv.FormatUint(jobID, 10))
			reportQ.ReleaseJob(jobID)
			continue
		}

		// Delete the job if it was successful
		reportQ.DeleteJob(jobID)
	}

}
