package wdmsapi

import (
	"context"
	"fmt"
	"net/http"
)

// CompanyService define available functions related to WDMS company
type CompanyService interface {
	List(context.Context, *CompanyListRequest) (*CompanyListResult, *Response, error)
	Get(context.Context, int) (*Company, *Response, error)
	Create(context.Context, *Company) (int, *Response, error)
	Update(context.Context, int, *Company) (*Response, error)
	Delete(context.Context, int) (*Response, error)
}

// Company represent company in WDMS
type Company struct {
	// surrogate key, given by system
	ID        int    `json:"id,omitempty"`
	CompanyID int    `json:"company_id"`
	Name      string `json:"company_name"`
}

// CompanyListRequest contains parameters for company search api
type CompanyListRequest struct {
	ListRequest

	// depart_id
	ID int `json:"depart_id,omitempty"`

	// depart_name
	Name string `json:"depart_name,omitempty"`

	// company
	CompanyID int `json:"company,omitempty"`

	// company_name
	CompanyName string `json:"company_name,omitempty"`
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
	path := "/api/companies/?" + queries.Encode()

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

func (s *companyService) Create(ctx context.Context, data *Company) (int, *Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodPost, "/api/companies/", data)
	if err != nil {
		return -1, nil, err
	}

	result := new(Company)
	resp, err := s.client.Do(ctx, r, result)
	if err != nil {
		return -1, nil, err
	}

	return result.ID, resp, nil
}

func (s *companyService) Update(ctx context.Context, id int, data *Company) (*Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("/api/companies/%d", id), data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, r, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *companyService) Delete(ctx context.Context, id int) (*Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/companies/%d", id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, r, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *companyService) Get(ctx context.Context, id int) (*Company, *Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/api/companies/%d", id), nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(Company)
	resp, err := s.client.Do(ctx, r, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp, nil
}
