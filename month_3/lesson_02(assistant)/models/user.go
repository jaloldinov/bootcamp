package models

type CreateUser struct {
	Login       string `json:"login"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Age         string `json:"age"`
	PhoneNumber string `json:"phone_number"`
}

type User struct {
	ID          string `json:"id"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Age         string `json:"age"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateUser struct {
	ID          string `json:"id"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Age         string `json:"age"`
	PhoneNumber string `json:"phone_number"`
}

type IdRequest struct {
	Id string `json:"id"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginDataRespond struct {
	Login       string `json:"login"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type LoginRespond struct {
	Token string
}

type PhoneNumberRequest struct {
	PhoneNumber string `json:"phone_number"`
}

type GetAllUserRequest struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Name  string `json:"name"`
}

type GetAllUser struct {
	Users []User `json:"Users"`
	Count int    `json:"count"`
}
