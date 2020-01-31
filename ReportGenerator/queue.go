package main

import (
	"log"
	"strconv"
	"time"

	"github.com/beanstalkd/go-beanstalk"
)

// GetJobFromQueue : Returns ready job from the queue
func GetJobFromQueue(tube *beanstalk.Conn) (id uint64, jobBody []byte, jobReady bool) {

	jobID, body, err := tube.Reserve(20 * time.Second)

	if err != nil {
		return 0, make([]byte, 0), false
	}

	return jobID, body, true
}

// ReleaseJob : Releases a job back to the queue
func ReleaseJob(tube *beanstalk.Conn, jobID uint64) {
	if err := tube.Release(jobID, 0, 0); err != nil {
		log.Println("Couldn't release job id: " + strconv.FormatUint(jobID, 10))
	}
}
