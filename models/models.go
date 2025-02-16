package models

type LoginRequest struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type SessionResponse struct {
	DID            string `json:"did"`
	Handle         string `json:"handle"`
	Email          string `json:"email"`
	EmailConfirmed bool   `json:"emailConfirmed"`
	AccessToken    string `json:"accessJwt"`
	RefreshToken   string `json:"refreshJwt"`
	Active         bool   `json:"active"`
}