package models

type RegisterRequest struct{}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRespond struct {
	Token string `json:"token"`
}

type RequestByUsername struct {
	Username string `json:"username"`
}
