package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ShareFrame/create-session/handler"
	"github.com/ShareFrame/create-session/models"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
)

type UnitEvent struct {
	Body string `json:"body"`
}

func handle(ctx context.Context, event UnitEvent) (models.SessionResponse, error) {
	var login models.LoginRequest
	if err := json.Unmarshal([]byte(event.Body), &login); err != nil {
		return models.SessionResponse{}, fmt.Errorf("invalid input: %w", err)
	}

	resp, err := handler.HandleLogin(ctx, login)
	if err != nil {
		return models.SessionResponse{}, err
	}
	
	logrus.WithFields(logrus.Fields{
		"sessionResp": fmt.Sprintf("%+v", resp),
	}).Info("Decoded session response")
	
	return *resp, nil
}

func main() {
	lambda.Start(handle)
}
