package main

import (
	"bytes"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/google/uuid"
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
func (s *S3) UploadCSVToS3(jobID uint64, userName string) (link string, ok bool, message string) {

	file, err := os.Open("./output/report" + strconv.FormatUint(jobID, 10) + ".csv")

	if err != nil {
		return "", false, "Unable to find CSV in File Path"
	}

	defer file.Close()

	objectKey := "/Reports/" + userName + "/report-" + uuid.New().String() + ".csv"

	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	file.Read(buffer)

	_, err = s3.New(s.session).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(s.Bucket),
		Key:           aws.String(objectKey),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(fileSize),
		ContentType:   aws.String(http.DetectContentType(buffer)),
	})

	if err != nil {
		return "", false, "Error while uploading file to S3"
	}

	return s.GetUploadedFileURL(objectKey), true, "Successfully uploaded to S3"
}

// GetUploadedFileURL : Generate download URL for the uploaded file
func (s *S3) GetUploadedFileURL(objectKey string) (fileURL string) {

	fileURL = "https://" + s.Bucket + ".s3." + s.Region + ".amazonaws.com" + objectKey
	return
}
