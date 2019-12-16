package wdmsapi

import (
	"context"
	"fmt"
	"net/http"
)

type IClockService interface {
	List(context.Context, *IClockListRequest) (*IClockListResult, *Response, error)
	Get(context.Context, int) (*IClock, *Response, error)
	Create(context.Context, *IClock) (*Response, error)
	Update(context.Context, int, *IClock) (*Response, error)
	Delete(context.Context, int) (*Response, error)
}

type IClock struct {
	ID               int    `json:"id,omitempty"`
	SN               string `json:"sn,omitempty"`
	Status           string `json:"status,omitempty"`
	LastActivity     Time   `json:"last_activity,omitempty"`
	Alias            string `json:"alias,omitempty"`
	IsMasterDevice   string `json:"is_master_device,omitempty"`
	FingerCount      string `json:"finger_count,omitempty"`
	TransactionCount string `json:"transaction_count,omitempty"`
	UserCount        string `json:"user_count,omitempty"`
	FaceCount        string `json:"face_count,omitempty"`
	PalmCount        string `json:"palm_count,omitempty"`
	DeviceName       string `json:"device_name,omitempty"`
	Area             int    `json:"area,omitempty"`
	CmdCount         string `json:"cmd_count,omitempty"`
	FWVersion        string `json:"fw_version,omitempty"`
	CompanyName      string `json:"company_name,omitempty"`
	IPAddress        string `json:"ip_address"`
}

type IClockListRequest struct {
	ListRequest

	SN    string `json:"sn,omitempty"`
	Alias string `json:"alias,omitempty"`
}

type IClockListResult struct {
	ListResult

	Data []IClock `json:"data"`
}

type iclockService struct {
	client *Client
}

func (s *iclockService) List(ctx context.Context, req *IClockListRequest) (*IClockListResult, *Response, error) {
	// ensure params not nil
	if req == nil {
		req = &IClockListRequest{}
	}

	// construct query string based on params
	queries := MarshalURLQuery(req)
	path := "/api/iclocks?" + queries.Encode()

	r, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(IClockListResult)
	resp, err := s.client.Do(ctx, r, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, err
}

func (s *iclockService) Create(ctx context.Context, data *IClock) (*Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodPost, "/api/iclocks", data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, r, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *iclockService) Update(ctx context.Context, id int, data *IClock) (*Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("/api/iclocks/%d", id), data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, r, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *iclockService) Delete(ctx context.Context, id int) (*Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/iclocks/%d", id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, r, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *iclockService) Get(ctx context.Context, id int) (*IClock, *Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/iclocks/%d", id), nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(IClock)
	resp, err := s.client.Do(ctx, r, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp, nil
}
