package main

// NotifierJob : Notifier Job Structure
type NotifierJob struct {
	User   User   `json:"user"`
	Search Search `json:"search"`
	Result Result `json:"result"`
}

// User : Contains User name and email
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Search : Contains information about the search hits
type Search struct {
	TotalHits int64 `json:"totalHits"`
}

// Result : Contains the final url for the report
type Result struct {
	URL string `json:"url"`
}
