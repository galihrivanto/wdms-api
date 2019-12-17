package wdmsapi

import "context"

type TransactionService interface {
	List(context.Context, *TransactionListRequest) (*TransactionListResult, *Response, error)
	Get(context.Context, int) (*Transaction, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

type Transaction struct {
	ID int `json:"id,omitempty"`
}

type TransactionListRequest struct {
	ListRequest

	Pin string `json:"pin,omitempty"`

	SN string `json:"sn,omitempty"`
}

type TransactionListResult struct {
	ListResult

	Data []Transaction `json:"data"`
}

// {id: 15103, time: "2019-12-17 18:24:52", status: "Check out", verify: "FP", workcode: "",…}
// id: 15103
// time: "2019-12-17 18:24:52"
// status: "Check out"
// verify: "FP"
// workcode: ""
// reserved: "0"
// createtime: "2019-12-17 10:24:53"
// employee: 149
// iclock: 11
// pin: "19"
// emp_name: "MIAH MOMEN"
// area_name: "soilbuild"
// depart_name: "5c872e27b2afd44868526fb5"
// sn: "CF2P191360102"
// att_photo: ""
