package main

import (
	"encoding/json"
	"os"
	"time"

	"gopkg.in/olivere/elastic.v3"
)

// ConvertTimeLayoutToISO : Convert UI form Layout to ISO Default
func ConvertTimeLayoutToISO(date string) (ISO time.Time, ok bool) {
	const UIFormat string = "01/02/2006"

	t, err := time.Parse(UIFormat, date)

	if err != nil {
		return time.Time{}, false
	}

	return t, true
}

// ParseESResultToModel : Convert the ES Hits Result to Flights Model
func ParseESResultToModel(hits []*elastic.SearchHit) (fs []Flight, ok bool) {
	var flights []Flight

	for _, hit := range hits {
		var flight Flight

		err := json.Unmarshal(*hit.Source, &flight)

		if err != nil {
			return flights, false
		}

		flights = append(flights, flight)
	}

	return flights, true
}

// EnsureENVSet : Makes sure that all the necessary ENV are set
func EnsureENVSet() (ok bool, missing string) {

	// Ports for queues and DB
	ReportQueuePort := os.Getenv("REPORT_QUEUE_PORT")
	NotifierQueuePort := os.Getenv("NOTIFIER_QUEUE_PORT")
	ESPort := os.Getenv("ES_PORT")

	// AWS
	S3Region := os.Getenv("S3_REGION")
	S3Bucket := os.Getenv("S3_BUCKET")
	AWSAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	AWSSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	if ReportQueuePort == "" {
		return false, "Report Queue Port"
	}

	if NotifierQueuePort == "" {
		return false, "Notifier Queue Port"
	}

	if ESPort == "" {
		return false, "Elasticsearch Port"
	}

	if S3Region == "" {
		return false, "S3 Region"
	}

	if S3Bucket == "" {
		return false, "S3 Bucket"
	}

	if AWSAccessKeyID == "" {
		return false, "AWS Access Key ID"
	}

	if AWSSecretAccessKey == "" {
		return false, "AWS Secret Access Key"
	}

	return true, ""
}
