package model

type AuthResponse struct {
	*User
	Token string `json:"token"`
}
