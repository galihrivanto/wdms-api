package wdmsapi

import (
	"context"
	"fmt"
	"net/http"
)

type (
	DeviceStatus       string
	MasterDeviceStatus string
)

var (
	DeviceActive   DeviceStatus = "1"
	DeviceInActive DeviceStatus = "0"

	MasterDevice MasterDeviceStatus = "111111111111"
	SlaveDevice  MasterDeviceStatus = "111000000000"
)

type IClockService interface {
	List(context.Context, *IClockListRequest) (*IClockListResult, *Response, error)
	Get(context.Context, int) (*IClock, *Response, error)
	Create(context.Context, *IClock) (*Response, error)
	Update(context.Context, int, *IClock) (*Response, error)
	Delete(context.Context, int) (*Response, error)
}

type IClock struct {
	ID               int                `json:"id,omitempty"`
	SN               string             `json:"sn,omitempty"`
	Status           DeviceStatus       `json:"status,omitempty"`
	LastActivity     Time               `json:"last_activity,omitempty"`
	Alias            string             `json:"alias,omitempty"`
	IsMasterDevice   MasterDeviceStatus `json:"is_master_device,omitempty"`
	FingerCount      Number             `json:"finger_count,omitempty"`
	TransactionCount Number             `json:"transaction_count,omitempty"`
	UserCount        Number             `json:"user_count,omitempty"`
	FaceCount        Number             `json:"face_count,omitempty"`
	PalmCount        Number             `json:"palm_count,omitempty"`
	DeviceName       string             `json:"device_name,omitempty"`
	Area             int                `json:"area,omitempty"`
	CmdCount         Number             `json:"cmd_count,omitempty"`
	FWVersion        string             `json:"fw_version,omitempty"`
	CompanyName      string             `json:"company_name,omitempty"`
	IPAddress        string             `json:"ip_address"`
}

func (i IClock) IsMaster() bool {
	return i.IsMasterDevice == MasterDevice
}

func (i IClock) IsOnline() bool {
	return i.Status == DeviceActive
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
	path := "/api/iclocks/?" + queries.Encode()

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
	r, err := s.client.NewRequest(ctx, http.MethodPost, "/api/iclocks/", data)
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
	r, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/api/iclocks/%d", id), nil)
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
