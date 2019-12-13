package wdmsapi

import (
	"testing"
	"time"
)

func TestMarshalURLQuery(t *testing.T) {
	params := struct {
		Value1 string    `json:"value1"`
		Value2 int       `json:"value2"`
		Value3 float64   `json:"value3"`
		Value4 time.Time `json:"value4"`
		Value5 bool      `json:"value5"`
		Value6 int       `json:"value6"`
		Value7 string    `json:"value7,omitempty"`
	}{
		Value1: "hello",
		Value2: 2,
		Value3: 3.14,
		Value4: time.Now(),
		Value5: true,
	}

	queries := MarshalURLQuery(params)
	for k, v := range queries {
		t.Log(k, v)
	}
}
