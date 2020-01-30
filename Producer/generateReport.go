package main

import (
	"encoding/json"
	"net/http"

	"strconv"
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
		ReturnErrorResponse(res, "Can not perform GET method on this endpoint", http.StatusMethodNotAllowed)
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
		ReturnErrorResponse(res, "Error while parsing form data. Please try again.", http.StatusInternalServerError)
		return
	}

	isError, message, jobID := PushJobToQueue(jobJSONData)

	if isError != false {
		ReturnErrorResponse(res, message, http.StatusInternalServerError)
		return
	}

	successResponse := SuccessResponse{
		Message:    "Successfully added your report to queue! Your report will be generated soon and sent to your email. Your Job ID is " + strconv.FormatUint(jobID, 10),
		Body:       reportJob,
		StatusCode: 200,
	}

	successJSONData, _ := json.Marshal(successResponse)

	res.Header().Set("Content-Type", "application/json")
	res.Write(successJSONData)

	return
}

// ReturnErrorResponse : Returns error response with custom message and statuscode
func ReturnErrorResponse(res http.ResponseWriter, message string, statusCode int) {
	res.WriteHeader(http.StatusInternalServerError)

	errorResponse := ErrorResponse{
		Message:    message,
		StatusCode: statusCode,
	}

	jsonData, _ := json.Marshal(errorResponse)

	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonData)
}
