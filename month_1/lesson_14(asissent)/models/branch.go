package models

type CreateBranch struct {
	Name   string
	Adress string
}

type Branch struct {
	Id     int
	Name   string
	Adress string
}

type IdRequest struct {
	Id int
}

type GetAllBranchRequest struct {
	Page  int
	Limit int
	Name  string
}

type GetAllBranch struct {
	Branches []Branch
	Count    int
}
