package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

const (
	// Sender is "From" address.
	Sender = "sender@example.com"

	// Recipient is "To" address.
	Recipient = "recipient@example.com"
	ConfigurationSet = "ConfigSet"
	AwsRegion = "us-west-2"
	Subject = "Amazon SES Test (AWS SDK for Go)"
	HTMLBody =  "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
	"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
	"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"

	TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."

	CharSet = "UTF-8"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:aws.String(AwsRegion)},
	)

	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{
			},
			ToAddresses: []*string{
				aws.String(Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data: aws.String(HTMLBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data: aws.String(TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data: aws.String(Subject),
			},
		},
		Source: aws.String(Sender),
		ConfigurationSetName: aws.String(ConfigurationSet),
	}

	result, err := svc.SendEmail(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
					fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
					fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
					fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
					fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println("Email Sent!")
	fmt.Println(result)
}

