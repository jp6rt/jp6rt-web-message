package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/sns"
)

var svc *sns.SNS

func init() {
	fmt.Println("init")
	region := "ap-southeast-1"
	session, err := session.NewSession(&aws.Config{
		Region: &region,
	})

	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to connect to AWS: %s", err.Error()))
		os.Exit(1)
	}

	svc = sns.New(session)
	fmt.Println("sns initiated")
}

// Handler ...
func Handler(ctx context.Context, e events.DynamoDBEvent) {
	var message string
	for _, record := range e.Records {

		// Print new values for attributes of type String
		for name, value := range record.Change.NewImage {
			if value.DataType() == events.DataTypeString && name == "message" {
				message = value.String()
			}
		}
	}

	topicARN := "arn:aws:sns:ap-southeast-1:230335894752:jp6rt-web-messages-dev-MessageTopic-Y6BDS59ZG771"

	result, err := svc.Publish(&sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: &topicARN,
	})

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(*result.MessageId)
}

func main() {
	lambda.Start(Handler)
}
