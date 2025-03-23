package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ShareFrame/create-session/models"
	"github.com/sirupsen/logrus"
)

func HandleLogin(ctx context.Context, event models.LoginRequest) (*models.SessionResponse, error) {
	logrus.WithFields(logrus.Fields{
		"identifier": event.Identifier,
	}).Info("Handling login request")

	baseURL := os.Getenv("ATPROTO_BASE_URL")
	blueskyAPI := fmt.Sprintf("%s/xrpc/com.atproto.server.createSession", baseURL)

	requestBody, err := json.Marshal(map[string]string{
		"identifier": event.Identifier,
		"password":   event.Password,
	})
	if err != nil {
		logrus.Error("Failed to marshal request body", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", blueskyAPI, bytes.NewBuffer(requestBody))
	if err != nil {
		logrus.Error("Failed to create HTTP request", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Error("Failed to execute HTTP request", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		logrus.WithFields(logrus.Fields{
			"status":   resp.StatusCode,
			"response": string(body),
		}).Error("Failed to authenticate with Bluesky API")
		return nil, fmt.Errorf("failed to authenticate with Bluesky API: %s", string(body))
	}

	var raw models.BskySessionResponse
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		logrus.Error("Failed to decode response body", err)
		return nil, err
	}

	sessionResp := models.SessionResponse{
		DID:            raw.DID,
		Handle:         raw.Handle,
		Email:          raw.Email,
		EmailConfirmed: raw.EmailConfirmed,
		AccessToken:    raw.AccessJwt,
		RefreshToken:   raw.RefreshJwt,
		Active:         raw.Active,
	}

	logrus.WithFields(logrus.Fields{
		"DID":    sessionResp.DID,
		"Handle": sessionResp.Handle,
	}).Info("Login successful")
	return &sessionResp, nil
}
