package wdmsapi

import (
	"context"
	"fmt"
	"net/http"
)

type EmployeeService interface {
	List(context.Context, *EmployeeListRequest) (*EmployeeListResult, *Response, error)
	Get(context.Context, int) (*Employee, *Response, error)
	Create(context.Context, *Employee) (*Response, error)
	Update(context.Context, int, *Employee) (*Response, error)
	Delete(context.Context, int) (*Response, error)
}

type Employee struct {
	ID               int    `json:"id,omitempty"`
	Pin              string `json:"pin,omitempty"`
	Name             string `json:"emp_name,omitempty"`
	Gender           string `json:"gender,omitempty"`
	Department       int    `json:"department,omitempty"`
	DepartmentName   string `json:"department_name,omitempty"`
	DepartmentID     string `json:"department_id,omitempty"`
	Card             string `json:"card,omitempty"`
	Priviledge       int    `json:"priviledge,omitempty"`
	MobileTel        string `json:"mobile_tel,omitempty"`
	FingerCount      string `json:"finger_count,omitempty"`
	FaceCount        string `json:"face_count,omitempty"`
	PalmCount        string `json:"palm_count,omitempty"`
	TransactionCount string `json:"transaction_count,omitempty"`
	Area             []int  `json:"area,omitempty"`
	AreaNames        string `json:"area_names,omitempty"`
	CompanyID        string `json:"company_id,omitempty"`
	CompanyName      string `json:"company_name,omitempty"`
}

type EmployeeListRequest struct {
	ListRequest

	Pin            string `json:"pin,omitempty"`
	EmployeeName   string `json:"emp_name,omitempty"`
	DepartmentID   string `json:"department,omitempty"`
	DepartmentName string `json:"depart_name,omitempty"`
}

type EmployeeListResult struct {
	ListResult

	Data []Employee `json:"data"`
}

type employeeService struct {
	client *Client
}

func (s *employeeService) List(ctx context.Context, req *EmployeeListRequest) (*EmployeeListResult, *Response, error) {
	if req == nil {
		req = &EmployeeListRequest{}
	}

	queries := MarshalURLQuery(req)
	path := "/api/employees?" + queries.Encode()

	r, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(EmployeeListResult)
	resp, err := s.client.Do(ctx, r, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, err
}

func (s *employeeService) Create(ctx context.Context, data *Employee) (*Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodPost, "/api/employees", data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, r, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *employeeService) Update(ctx context.Context, id int, data *Employee) (*Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("/api/employees/%d", id), data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, r, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *employeeService) Delete(ctx context.Context, id int) (*Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/employees/%d", id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, r, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *employeeService) Get(ctx context.Context, id int) (*Employee, *Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/employees/%d", id), nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(Employee)
	resp, err := s.client.Do(ctx, r, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp, nil
}
