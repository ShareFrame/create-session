package models

type LoginRequest struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type BskySessionResponse struct {
	DID            string `json:"did"`
	Handle         string `json:"handle"`
	Email          string `json:"email"`
	EmailConfirmed bool   `json:"emailConfirmed"`
	AccessJwt      string `json:"accessJwt"`
	RefreshJwt     string `json:"refreshJwt"`
	Active         bool   `json:"active"`
}

type SessionResponse struct {
	DID            string `json:"did"`
	Handle         string `json:"handle"`
	Email          string `json:"email"`
	EmailConfirmed bool   `json:"emailConfirmed"`
	AccessToken    string `json:"accessToken"`
	RefreshToken   string `json:"refreshToken"`
	Active         bool   `json:"active"`
}
