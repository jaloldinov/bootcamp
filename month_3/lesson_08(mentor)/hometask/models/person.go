package models

type CreatePerson struct {
	Name     string `json:"name"`
	Job      string `json:"job"`
	Age      int    `json:"age"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type ChangePassword struct {
	NewPassword string `json:"name"`
	OldPassword string `json:"password"`
}
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginRes struct {
	Token string `json:"token"`
}
type Person struct {
	Id        string `json:"id"`
	BranchId  string `json:"branch_id"`
	Name      string `json:"name"`
	Job       string `json:"job"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
}

type RequestByID struct {
	ID string
}
type RequestByUsername struct {
	Username string
}
type GetAllPersonsResponse struct {
	Persons []Person `json:"persons"`
	Count   int32    `json:"count"`
}
type GetAllPersonsRequest struct {
	Search string
	Job    string
	Age    int
	Page   int
	Limit  int
}
