package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler function Using AWS Lambda Proxy Request
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	switch request.HTTPMethod {
	case http.MethodGet:
		return getRequest(request)
	case http.MethodPost:
		return postRequest(request)
	default:
		return notImplemented()
	}
}

func getRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.PathParameters["name"]
	message := fmt.Sprintf(" { \"Message\" : \"Hello %s \" } ", name)
	return events.APIGatewayProxyResponse{Body: message, StatusCode: 200}, nil
}

func postRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.PathParameters["name"]
	message := fmt.Sprintf(" { \"Message\" : \"Creating a new user as %s \" } ", name)
	return events.APIGatewayProxyResponse{Body: message, StatusCode: 201}, nil
}

func notImplemented() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: "Not Implemented", StatusCode: http.StatusNotImplemented}, nil
}

func main() {
	lambda.Start(Handler)
}
