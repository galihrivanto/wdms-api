package wdmsapi

type Area struct {
	ID        string `json:"area_id"`
	Name      string `json:"area_name"`
	CompanyID string `json:"company"`
}

type AreaListRequest struct {
	ListRequest

	AreaID     string `json:"area_id"`
	AreaName   string `json:"area_name"`
	Company    string `json:"company"`
	Department string `json:"department"`
}

type AreaListResult struct {
	ListResult

	Data []Area
}
