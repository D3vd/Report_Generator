package main

import (
	"bytes"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3 structure
type S3 struct {
	session *session.Session

	Region string
	Bucket string
}

// Init : Initialize S3 instance
func (s *S3) Init() (ok bool) {

	s.Region = os.Getenv("S3_REGION")
	s.Bucket = os.Getenv("S3_BUCKET")

	if s.Region == "" || s.Bucket == "" {
		return false
	}

	session, err := session.NewSession(&aws.Config{Region: aws.String(s.Region)})

	if err != nil {
		return false
	}

	s.session = session

	return true
}

// UploadCSVToS3 : Upload the CSV file to S3
func (s *S3) UploadCSVToS3(filePath string) (link string, ok bool, message string) {

	file, err := os.Open(filePath)

	if err != nil {
		return "", false, "Unable to find CSV in File Path"
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	file.Read(buffer)

	_, err = s3.New(s.session).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(s.Bucket),
		Key:           aws.String(filePath),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(fileSize),
		ContentType:   aws.String(http.DetectContentType(buffer)),
	})

	if err != nil {
		return "", false, "Error while uploading file to S3"
	}

	return "", true, "Successfully uploaded to S3"
}
