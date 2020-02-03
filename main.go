package main

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	s3utils "./s3utils"
)

var bucketName = "<bucket-name>"

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	fmt.Println(err.Error())
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
		Body: fmt.Sprintf("%v", err.Error()),
	}, nil
}

// lambda on start execute function
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	pathParameters := request.PathParameters

	path1 := pathParameters["path1"]
	filename := pathParameters["filename"]

	// provide a path where file will be stored in s3
	key := fmt.Sprintf("folderName/%s/%s", path1, filename)
	body, err := base64.StdEncoding.DecodeString(request.Body)

	// upload file to s3
	err = s3utils.Upload(bucketName, key, body)

	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type":                "application/octet-stream",
			"Access-Control-Allow-Origin": "*",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
