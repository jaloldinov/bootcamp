package models

type RegisterRequest struct{}

type LoginRequest struct {
	Login    string
	Password string
}
