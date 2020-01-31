package main

import (
	"log"
	"strconv"
	"time"

	"github.com/beanstalkd/go-beanstalk"
)

// Queue structure
type Queue struct {
	tube *beanstalk.Conn
}

// Init Beanstalk Queue Connection
func (q *Queue) Init(port string) (ok bool) {
	tube, err := beanstalk.Dial("tcp", port)

	if err != nil {
		return false
	}

	q.tube = tube
	return true
}

// GetJobFromQueue : Returns ready job from the queue
func (q *Queue) GetJobFromQueue() (id uint64, jobBody []byte, jobReady bool) {

	jobID, body, err := q.tube.Reserve(20 * time.Second)

	if err != nil {
		return 0, make([]byte, 0), false
	}

	return jobID, body, true
}

// ReleaseJob : Releases a job back to the queue
func (q *Queue) ReleaseJob(jobID uint64) {
	if err := q.tube.Release(jobID, 0, 0); err != nil {
		log.Println("Couldn't release job id: " + strconv.FormatUint(jobID, 10), err)
	}
}


// CloseQueue : Close beanstalk connection
func (q *Queue) CloseQueue() {
	if err := q.tube.Close(); err != nil {
		log.Println("Error while closing queue Connection", err)
	}
}
