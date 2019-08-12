package main

import "github.com/aws/aws-lambda-go/lambda"

type Message struct {
	CreatedAt string
	Message   string
}

type response struct {
	Message string `json:"message"`
}

func handler() (response, error) {
	return response{
		Message: "Okay so your other function also executed successfully!",
	}, nil
}

func main() {
	lambda.Start(handler)
}
