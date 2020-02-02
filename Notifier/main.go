package main

import (
	"encoding/json"
	"log"
	"strconv"
)

func main() {

	NotifierQueuePort := "127.0.0.1:11300"

	var notifierQ Queue

	// Connect to thw Notifier Queue
	if ok := notifierQ.Init(NotifierQueuePort); !ok {
		log.Fatalln("Unable to connect to Notifier Queue at port " + NotifierQueuePort + ". Ensure it is active.")
		return
	}
	defer notifierQ.CloseQueue()

	for {

		jobID, jobBody, jobReady := notifierQ.GetJobFromQueue()

		if !jobReady {
			continue
		}

		var notifierJob NotifierJob

		if err := json.Unmarshal(jobBody, &notifierJob); err != nil {
			log.Println("Error while parsing notifier Job ID : " + strconv.FormatUint(jobID, 10))
			notifierQ.BuryJob(jobID)
			continue
		}

		if ok := SendEmailToUser(notifierJob); !ok {
			log.Println("Error while sending email.")
			notifierQ.ReleaseJob(jobID)
			continue
		}

		notifierQ.DeleteJob(jobID)
	}
}
