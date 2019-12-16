package wdmsapi

import (
	"context"
	"fmt"
	"net/http"
)

type DepartmentService interface {
	List(context.Context, *DepartmentListRequest) (*DepartmentListResult, *Response, error)
	Get(context.Context, int) (*Department, *Response, error)
	Create(context.Context, *Department) (*Response, error)
	Update(context.Context, int, *Department) (*Response, error)
	Delete(context.Context, int) (*Response, error)
}

type Department struct {
	ID               int    `json:"depart_id,omitempty"`
	Name             string `json:"depart_name,omitempty"`
	CompanyID        int    `json:"company,omitempty"`
	ParentDepartment int    `json:"parent_depart,omitempty"`
}

type DepartmentListRequest struct {
	ListRequest

	DepartmentID   int    `json:"depart_id,omitempty"`
	DepartmentName string `json:"depart_name,omitempty"`
	Company        int    `json:"company,omitempty"`
}

type DepartmentListResult struct {
	ListResult

	Data []Department `json:"data"`
}

type departmentService struct {
	client *Client
}

func (s *departmentService) List(ctx context.Context, req *DepartmentListRequest) (*DepartmentListResult, *Response, error) {
	// ensure params not nil
	if req == nil {
		req = &DepartmentListRequest{}
	}

	// construct query string based on params
	queries := MarshalURLQuery(req)
	path := "/api/departments?" + queries.Encode()

	r, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(DepartmentListResult)
	resp, err := s.client.Do(ctx, r, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, err
}

func (s *departmentService) Create(ctx context.Context, data *Department) (*Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodPost, "/api/departments", data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, r, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *departmentService) Update(ctx context.Context, id int, data *Department) (*Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("/api/departments/%d", id), data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, r, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *departmentService) Delete(ctx context.Context, id int) (*Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/departments/%d", id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, r, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *departmentService) Get(ctx context.Context, id int) (*Department, *Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/departments/%d", id), nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(Department)
	resp, err := s.client.Do(ctx, r, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp, nil
}
