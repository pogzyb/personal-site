package website

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"log"
	"os"
)

var sess *session.Session

func initSession() {
	var err error
	sess, err = session.NewSession(&aws.Config{})
	if err != nil {
		log.Fatalf("could not create aws session: %v", err)
	}
}

func sendEmail(subject string, serverError error, stacktrace string) {
	svc := ses.New(sess)
	body := fmt.Sprintf("<p>Fatal Error!</p><p>error:<br>%v</p><p><br>%s</p>", serverError, stacktrace)
	email := ses.SendEmailInput{
		Source: aws.String("alerts@joeobarzanek.com"),
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(os.Getenv("EMAIL_500")),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
		},
	}
	_, err := svc.SendEmail(&email); if err != nil {
		log.Printf("could not send email with SES: %v", err)
	}
}
