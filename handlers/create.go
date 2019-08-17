package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response ..
type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	x := Response{
		Message: "ok",
		Code:    200,
	}
	response, err := json.Marshal(x)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 500}, nil
}

func main() {
	lambda.Start(handler)
}
