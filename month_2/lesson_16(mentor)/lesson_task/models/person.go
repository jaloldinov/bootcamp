package models

type CreatePerson struct {
	Name string
	Job  string
	Age  int
}
type Person struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Job  string `json:"job"`
	Age  int    `json:"age"`
}

type RequestByID struct {
	ID string
}

type GetAllRequest struct {
	Page  int
	Limit int
}
