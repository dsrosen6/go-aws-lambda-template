package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// The data that is passed to the Lambda function
type Event struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// The data that is returned from the Lambda function
type Response struct {
	Message string `json:"answer"`
}

// Entry point, no action needed
func main() {
	lambda.Start(HandleLambdaEvent)
}

// Lambda function handler - main logic starts here
func HandleLambdaEvent(event *Event) (*Response, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}
	return &Response{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}
