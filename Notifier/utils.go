package main

import "os"

// EnsureENVSet : Ensure that all the required variables are set
func EnsureENVSet() (ok bool) {
	NotifierPort := os.Getenv("NOTIFIER_QUEUE_PORT")
	SendGridAPIKey := os.Getenv("SENDGRID_API_KEY")

	if NotifierPort == "" || SendGridAPIKey == "" {
		return false
	}

	return true
}
