package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

func main() {

	if ok := EnsureENVSet(); !ok {
		log.Fatalln("Environment variables are not set. Have to set variables before continuing.")
		return
	}

	NotifierQueuePort := os.Getenv("NOTIFIER_QUEUE_PORT")

	var notifierQ Queue

	// Connect to thw Notifier Queue
	if ok := notifierQ.Init(NotifierQueuePort); !ok {
		log.Fatalln("Unable to connect to Notifier Queue at port " + NotifierQueuePort + ". Ensure it is active.")
		return
	}
	defer notifierQ.CloseQueue()

	for {

		// Get job from the notifier queue
		jobID, jobBody, jobReady := notifierQ.GetJobFromQueue()

		// If the job is not ready then try again
		if !jobReady {
			continue
		}

		log.Println("Reserving Job " + strconv.FormatUint(jobID, 10) + " from Notifier Queue.")

		// Unmarshal notifier job
		var notifierJob NotifierJob

		if err := json.Unmarshal(jobBody, &notifierJob); err != nil {
			log.Println("Error while parsing notifier Job ID : " + strconv.FormatUint(jobID, 10))
			notifierQ.BuryJob(jobID)
			continue
		}

		// Check if the total hits is zero
		if notifierJob.Search.TotalHits == 0 {
			log.Println("Job " + strconv.FormatUint(jobID, 10) + " has no hits. Sending unsuccessful email to user.")
			if ok := SendUnsuccessEmailToUser(notifierJob); !ok {
				log.Println("Error while sending email.")
				notifierQ.ReleaseJob(jobID)
				continue
			}
			notifierQ.DeleteJob(jobID)
			continue
		}

		log.Println("Sending Successful Email to User for Job: " + strconv.FormatUint(jobID, 10))
		// Send email to the user
		if ok := SendSuccessEmailToUser(notifierJob); !ok {
			log.Println("Error while sending email.")
			notifierQ.ReleaseJob(jobID)
			continue
		}

		// Delete the job from the queue after it has been successfully completed
		notifierQ.DeleteJob(jobID)
	}
}
