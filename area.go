package wdmsapi

import (
	"context"
	"fmt"
	"net/http"
)

type AreaService interface {
	List(context.Context, *AreaListRequest) (*AreaListResult, *Response, error)
	Get(context.Context, int) (*Area, *Response, error)
	Create(context.Context, *Area) (*Response, error)
	Update(context.Context, int, *Area) (*Response, error)
	Delete(context.Context, int) (*Response, error)
}

type Area struct {
	ID        int    `json:"area_id,omitempty"`
	Name      string `json:"area_name,omitempty"`
	CompanyID int    `json:"company,omitempty"`
}

type AreaListRequest struct {
	ListRequest

	AreaID     int    `json:"area_id,omitempty"`
	AreaName   string `json:"area_name,omitempty"`
	Company    int    `json:"company,omitempty"`
	Department string `json:"department,omitempty"`
}

type AreaListResult struct {
	ListResult

	Data []Area
}

type areaService struct {
	client *Client
}

func (s *areaService) List(ctx context.Context, req *AreaListRequest) (*AreaListResult, *Response, error) {
	if req == nil {
		req = &AreaListRequest{}
	}

	queries := MarshalURLQuery(req)
	path := "/api/areas?" + queries.Encode()

	r, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(AreaListResult)
	resp, err := s.client.Do(ctx, r, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, err
}

func (s *areaService) Create(ctx context.Context, data *Area) (*Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodPost, "/api/areas", data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, r, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *areaService) Update(ctx context.Context, id int, data *Area) (*Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("/api/areas/%d", id), data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, r, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *areaService) Delete(ctx context.Context, id int) (*Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/areas/%d", id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, r, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *areaService) Get(ctx context.Context, id int) (*Area, *Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/areas/%d", id), nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(Area)
	resp, err := s.client.Do(ctx, r, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp, nil
}
