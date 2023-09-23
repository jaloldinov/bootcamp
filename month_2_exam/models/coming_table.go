package models

type ComingTablePrimaryKey struct {
	Id string `json:"id"`
}

type ComingTableComingIdKey struct {
	ComingId string `json:"coming_id"`
}

type CreateComingTable struct {
	ComingId string `json:"coming_id"`
	BranchId string `json:"branch_id"`
}

type ComingTable struct {
	Id        string `json:"id"`
	ComingId  string `json:"coming_id"`
	BranchId  string `json:"branch_id"`
	DateTime  string `json:"date_time"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateComingTable struct {
	Id       string `json:"id"`
	ComingId string `json:"coming_id"`
	BranchId string `json:"branch_id"`
	DateTime string `json:"date_time"`
}

type ComingTableGetListRequest struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type ComingTableGetListResponse struct {
	Count        int            `json:"count"`
	ComingTables []*ComingTable `json:"products"`
}
