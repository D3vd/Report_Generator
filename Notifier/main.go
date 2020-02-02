package main

import "log"

func main() {

	NotifierQueuePort := "127.0.0.1:11300"

	var notifierQ Queue

	if ok := notifierQ.Init(NotifierQueuePort); !ok {
		log.Fatalln("Unable to connect to Notifier Queue at port " + NotifierQueuePort + ". Ensure it is active.")
		return
	}
	defer notifierQ.CloseQueue()

}
