package models

// ResponseModel ...
type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type ErrorResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Status struct {
	Status string `json:"status"`
}

type LoginReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRes struct {
	Token string `json:"token"`
}
