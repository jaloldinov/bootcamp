package models

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type CreateUser struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Balance int    `json:"balance"`
}

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Balance int    `json:"balance"`
}

type UpdateUser struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Balance int    `json:"balance"`
}

type UserGetListRequest struct {
	Offset int
	Limit  int
}

type UserGetListResponse struct {
	Count int
	Users []*User
}

type Task3Model struct {
	Name  string
	Price int
	Count int
	Total int
	Time  string
}
