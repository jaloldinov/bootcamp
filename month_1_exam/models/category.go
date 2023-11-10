package models

type CategoryPrimaryKey struct {
	Id string `json:"id"`
}

type CreateCategory struct {
	Name string `json:"name"`
}

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UpdateCategory struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CategoryGetListRequest struct {
	Offset int
	Limit  int
}

type CategoryGetListResponse struct {
	Count     int
	Categorys []*Category
}
