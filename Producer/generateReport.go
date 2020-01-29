package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// SuccessResponse : Generate success response
type SuccessResponse struct {
	Message    string    `json:"message"`
	Body       ReportJob `json:"body"`
	StatusCode int       `json:"status_code"`
}

// GenerateReportJob ...
func GenerateReportJob(res http.ResponseWriter, req *http.Request) {

	// Check if the request method is POST
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(res, "invalid_http_method")
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

	jsonData, err := json.Marshal(reportJob)

	if err != nil {
		log.Println("Error while parsing form data")
		fmt.Fprintf(res, "Error while parsing form data. Please try again.")
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonData)

	return
}
