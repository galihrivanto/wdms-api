package wdmsapi

type Department struct {
	ID               string `json:"depart_id"`
	Name             string `json:"depart_name"`
	CompanyID        string `json:"company"`
	ParentDepartment string `json:"parent_depart"`
}

type DepartmentListRequest struct {
	ListRequest

	DepartmentID   string `json:"area_id"`
	DepartmentName string `json:"area_name"`
	Company        string `json:"company"`
}

type DepartmentListResult struct {
	ListResult

	Data []Department `json:"data"`
}
