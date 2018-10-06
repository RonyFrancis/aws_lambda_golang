// main.go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"math"
	"net/http"
)

// Response json
type Response struct {
	Key         int `json:"key"`
	SquareValue int `json:"square_value"`
}

// Request json
type Request struct {
	Key int `json:"key"`
}

// hello lambda function gives back square of the key got from the request
func hello(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var request Request
	if err := json.Unmarshal([]byte(req.Body), &request); err != nil {
		fmt.Println(err)
	}
	fmt.Println(request)
	res := Response{Key: request.Key, SquareValue: request.Key * request.Key}
	js, _ := json.Marshal(res)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
