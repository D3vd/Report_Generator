package main

// ReportJob : Final structure of the job that is pushed to the queue
type ReportJob struct {
	UserInfo  UserInfo  `json:"user"`
	QueryBody QueryBody `json:"query"`
}

// UserInfo : User Name and Email
type UserInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// QueryBody : Contains the report instructions
type QueryBody struct {
	CarrierName string `json:"carrier"`
	StartDate   string `json:"start"`
	EndDate     string `json:"end"`
}
