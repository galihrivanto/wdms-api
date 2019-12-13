package wdmsapi

import (
	"context"
	"net/http"
)

// CompanyService define available functions related to WDMS company
type CompanyService interface {
	List(context.Context, *CompanyListRequest) (*CompanyListResult, *Response, error)
}

// Company represent company in WDMS
type Company struct {
	ID   string `json:"company_id"`
	Name string `json:"company_name"`
}

// CompanyListRequest contains parameters for company search api
type CompanyListRequest struct {
	ListRequest

	// depart_id
	ID string `json:"depart_id,omitempty"`

	// depart_name
	Name string `json:"depart_name,omitempty"`

	// company
	CompanyID string `json:"company,omitempty"`

	// company_name
	CompanyName string `json:"company_name,omitempty"`

	// company_id_icontains
	CompanyIDContains string `json:"company_id_icontains,omitempty"`

	// depart_id_icontains
	DepartmentIDContains string `json:"depart_id_icontains,omitempty"`

	// depart_name_icontains
	DepartmentNameContains string `json:"depart_name_icontains,omitempty"`
}

type CompanyListResult struct {
	ListResult

	Data []Company `json:"data"`
}

type companyService struct {
	client *Client
}

func (s *companyService) List(ctx context.Context, req *CompanyListRequest) (*CompanyListResult, *Response, error) {
	// ensure params not nil
	if req == nil {
		req = &CompanyListRequest{}
	}

	// construct query string based on params
	queries := MarshalURLQuery(req)
	path := "/api/companies?" + queries.Encode()

	r, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(CompanyListResult)
	resp, err := s.client.Do(ctx, r, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, err
}
