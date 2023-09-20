package models

type BranchPrimaryKey struct {
	Id string `json:"id"`
}
type CreateBranch struct {
	Name string `json:"name"`
}

type Branch struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type BranchGetListRequest struct {
	Offset int
	Limit  int
}

type BranchGetListResponse struct {
	Count    int
	Branches []*Branch
}
