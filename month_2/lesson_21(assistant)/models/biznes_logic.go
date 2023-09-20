package models

type TopStaff struct {
	Name      string  `json:"name"`
	Branch    string  `json:"branch"`
	Total_Sum float64 `json:"total_sum"`
	StaffType string  `json:"staff_type"`
}

type TopStaffRequest struct {
	FromDate string `json:"from_date"`
	ToDate   string `json:"to_date"`
}

type TopStaffResponse struct {
	TopStaffs []*TopStaff
}
