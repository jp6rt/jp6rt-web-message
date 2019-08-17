package main

import "github.com/aws/aws-lambda-go/lambda"

type response struct {
	Message string `json:"message"`
}

func handler() (response, error) {
	return response{
		Message: "test message",
	}, nil
}

func main() {
	lambda.Start(handler)
}
