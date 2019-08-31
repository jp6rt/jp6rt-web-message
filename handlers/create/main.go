package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	uuid "github.com/satori/go.uuid"
)

// Message ..
type Message struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	CreatedAt string `json:"createdAt"`
}

var db *dynamodb.DynamoDB

func init() {
	fmt.Println("init")
	region := os.Getenv("region")
	session, err := session.NewSession(&aws.Config{
		Region: &region,
	})

	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to connect to AWS: %s", err.Error()))
		os.Exit(1)
	}

	db = dynamodb.New(session)
	fmt.Println("db initiated")
}

// Handler ..
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Handler")

	uuid := uuid.Must(uuid.NewV1(), nil).String()
	tableName := aws.String(os.Getenv("messageTable"))
	now := time.Now()

	message := &Message{
		ID:        uuid,
		CreatedAt: now.String(),
	}

	headers := map[string]string{
		"Access-Control-Allow-Origin": "*",
		"Content-Type":                "application/json",
	}

	json.Unmarshal([]byte(request.Body), message)

	item, _ := dynamodbattribute.MarshalMap(message)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: tableName,
	}

	_, err := db.PutItem(input)

	if err != nil {
		fmt.Println("Error putting item")
		return events.APIGatewayProxyResponse{
			Headers:    headers,
			Body:       err.Error(),
			StatusCode: 500,
		}, nil
	}

	body, _ := json.Marshal(message)
	return events.APIGatewayProxyResponse{
		Headers:    headers,
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
