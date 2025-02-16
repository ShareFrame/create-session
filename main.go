package main

import (
	"github.com/ShareFrame/create-session/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.HandleLogin)
}