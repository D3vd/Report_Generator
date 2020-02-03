package main

import (
	"os"
	"strconv"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendSuccessEmailToUser : Send an email to the user for Successful Report Generation using sendgrid API
func SendSuccessEmailToUser(notifierJob NotifierJob) (ok bool) {

	from := mail.NewEmail("Report Generator", "report.generator@gmail.com")
	to := mail.NewEmail(notifierJob.User.Name, notifierJob.User.Email)

	subject := "Your Flights Report has been successfully Generated!"

	htmlContent := `
		Hey ` + notifierJob.User.Name + `! <br>
		<br>
		Your report was successfully generated. It can be found <a href="` + notifierJob.Result.URL + `">here</a>. <br>
		Got <strong>` + strconv.FormatInt(notifierJob.Search.TotalHits, 10) + `</strong> hits for your query. <br>
    <br>
		Thanks!
	`

	message := mail.NewSingleEmail(from, subject, to, "Flights Report", htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	_, err := client.Send(message)

	if err != nil {
		return false
	}

	return true
}

// SendUnsuccessEmailToUser : Sends unsuccess email to user if query hits is zero
func SendUnsuccessEmailToUser(notifierJob NotifierJob) (ok bool) {

	from := mail.NewEmail("Report Generator", "report.generator@gmail.com")
	to := mail.NewEmail(notifierJob.User.Name, notifierJob.User.Email)

	subject := "We didn't get any hits for your Query"

	htmlContent := `
		Hey ` + notifierJob.User.Name + `! <br>
		<br>
		Your query did not yield any results from our database... Please try widening your query. <br>
    <br>
		Thanks!
	`

	message := mail.NewSingleEmail(from, subject, to, "Flights Report", htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	_, err := client.Send(message)

	if err != nil {
		return false
	}

	return true
}
