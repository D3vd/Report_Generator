package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SuccessResponse : Generate success response
type SuccessResponse struct {
	Message    string    `json:"message"`
	Body       ReportJob `json:"body"`
	StatusCode int       `json:"status_code"`
}

// ErrorResponse : Generate error response
type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// GenerateReportJob ...
func GenerateReportJob(res http.ResponseWriter, req *http.Request) {

	// Check if the request method is POST
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusMethodNotAllowed)

		errorResponse := ErrorResponse{
			Message:    "Can not perform GET method on this endpoint",
			StatusCode: 405,
		}

		jsonData, _ := json.Marshal(errorResponse)

		res.Header().Set("Content-Type", "application/json")
		res.Write(jsonData)

		return
	}

	req.ParseForm()

	userInfo := UserInfo{
		Name:  req.Form.Get("userName"),
		Email: req.Form.Get("userEmail"),
	}

	queryBody := QueryBody{
		CarrierName: req.Form.Get("carrier"),
		StartDate:   req.Form.Get("start"),
		EndDate:     req.Form.Get("end"),
	}

	reportJob := ReportJob{
		UserInfo:  userInfo,
		QueryBody: queryBody,
	}

	jobJSONData, err := json.Marshal(reportJob)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)

		errorResponse := ErrorResponse{
			Message:    "Error while parsing form data. Please try again.",
			StatusCode: 500,
		}

		jsonData, _ := json.Marshal(errorResponse)

		res.Header().Set("Content-Type", "application/json")
		res.Write(jsonData)

		return
	}

	// TODO: Add function to push the report job to the queue
	fmt.Println(string(jobJSONData))

	successResponse := SuccessResponse{
		Message:    "Successfully added your report to queue! Your report will be generated soon and sent to your email.",
		Body:       reportJob,
		StatusCode: 200,
	}

	successJSONData, _ := json.Marshal(successResponse)

	res.Header().Set("Content-Type", "application/json")
	res.Write(successJSONData)

	return
}
