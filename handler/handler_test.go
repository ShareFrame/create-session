package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ShareFrame/create-session/models"
	"github.com/stretchr/testify/assert"
)

func TestHandleLogin(t *testing.T) {
	os.Setenv("ATPROTO_BASE_URL", "http://mockapi.com")

	tests := []struct {
		name           string
		request        models.LoginRequest
		mockResponse   string
		mockStatusCode int
		expectError    bool
		expectResponse *models.SessionResponse
	}{
		{
			name: "Successful login",
			request: models.LoginRequest{
				Identifier: "user@example.com",
				Password:   "password123",
			},
			mockResponse: `{"did":"did:example:123","handle":"user@example.com"}`,
			mockStatusCode: http.StatusOK,
			expectError:    false,
			expectResponse: &models.SessionResponse{
				DID:    "did:example:123",
				Handle: "user@example.com",
			},
		},
		{
			name: "Invalid credentials",
			request: models.LoginRequest{
				Identifier: "user@example.com",
				Password:   "wrongpassword",
			},
			mockResponse:   `{"error":"Invalid credentials"}`,
			mockStatusCode: http.StatusUnauthorized,
			expectError:    true,
			expectResponse: nil,
		},
		{
			name: "Server error",
			request: models.LoginRequest{
				Identifier: "user@example.com",
				Password:   "password123",
			},
			mockResponse:   `{"error":"Internal server error"}`,
			mockStatusCode: http.StatusInternalServerError,
			expectError:    true,
			expectResponse: nil,
		},
		{
			name: "Malformed response body",
			request: models.LoginRequest{
				Identifier: "user@example.com",
				Password:   "password123",
			},
			mockResponse:   `{"did":123, "handle":}`,
			mockStatusCode: http.StatusOK,
			expectError:    true,
			expectResponse: nil,
		},
		{
			name: "Empty response body",
			request: models.LoginRequest{
				Identifier: "user@example.com",
				Password:   "password123",
			},
			mockResponse:   ``,
			mockStatusCode: http.StatusOK,
			expectError:    true,
			expectResponse: nil,
		},
		{
			name: "Failed to create HTTP request",
			request: models.LoginRequest{
				Identifier: "user@example.com",
				Password:   "password123",
			},
			mockResponse:   "",
			mockStatusCode: http.StatusOK,
			expectError:    true,
			expectResponse: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.mockStatusCode)
				w.Write([]byte(tt.mockResponse))
			}))
			defer server.Close()

			os.Setenv("ATPROTO_BASE_URL", server.URL)

			response, err := HandleLogin(context.Background(), tt.request)
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, response)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, tt.expectResponse, response)
			}
		})
	}
}
