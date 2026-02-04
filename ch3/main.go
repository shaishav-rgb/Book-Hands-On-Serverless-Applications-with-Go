package main

import "github.com/aws/aws-lambda-go/lambda"

func handler() (string, error) {
	return "Welcome to Serverless world uploaded from s3", nil
}

func main() {
	lambda.Start(handler)
}
