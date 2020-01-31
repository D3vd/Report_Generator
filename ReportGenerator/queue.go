package main

import (
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
