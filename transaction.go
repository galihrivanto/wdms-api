package wdmsapi

import (
	"context"
	"fmt"
	"net/http"
)

type VerifyKind string

var (
	VerifyByFace        = "Face"
	VerifyByFingerPrint = "FP"
	VerifyByCard        = "Card"
)

type TransactionService interface {
	List(context.Context, *TransactionListRequest) (*TransactionListResult, *Response, error)
	Get(context.Context, int) (*Transaction, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

type Transaction struct {
	ID int `json:"id,omitempty"`

	// worker's pin
	Pin string `json:"pin,omitempty"`

	// device's number
	SN string `json:"sn,omitempty"`

	Verify         VerifyKind `json:"verify"`
	WorkCode       string     `json:"workcode"`
	Reserved       Number     `json:"reserved"`
	Employee       int        `json:"employee"`
	IClock         int        `json:"iclock"`
	EmployeeName   string     `json:"emp_name"`
	AreaName       string     `json:"area_name"`
	DepartmentName string     `json:"depart_name"`
	AttPhoto       string     `json:"att_photo"`
	Status         string     `json:"status"`
	Time           Time       `json:"time"`
	CreateTime     Time       `json:"createtime"`
}

type TransactionListRequest struct {
	ListRequest

	// worker's pin
	Pin string `json:"pin,omitempty"`

	// device's number
	SN string `json:"sn,omitempty"`

	StartDate Time `json:"start_date"`
	EndDate   Time `json:"end_date"`
}

type TransactionListResult struct {
	ListResult

	Data []Transaction `json:"data"`
}

type transactionService struct {
	client *Client
}

func (s *transactionService) List(ctx context.Context, req *TransactionListRequest) (*TransactionListResult, *Response, error) {
	// ensure params not nil
	if req == nil {
		req = &TransactionListRequest{}
	}

	// construct query string based on params
	queries := MarshalURLQuery(req)
	path := "/api/transactions/?" + queries.Encode()

	r, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(TransactionListResult)
	resp, err := s.client.Do(ctx, r, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, err
}

func (s *transactionService) Delete(ctx context.Context, id int) (*Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/transactions/%d", id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, r, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *transactionService) Get(ctx context.Context, id int) (*Transaction, *Response, error) {
	r, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/api/transactions/%d", id), nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(Transaction)
	resp, err := s.client.Do(ctx, r, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp, nil
}
