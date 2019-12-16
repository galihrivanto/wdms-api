package wdmsapi

import (
	"strconv"
	"time"
)

const timeFormat = "2006-01-02 15:04"

type Time struct {
	time.Time
}

func (t *Time) String() string {
	return t.Time.String()
}

func (t *Time) UnmarshalJSON(data []byte) error {
	str, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	t.Time, err = time.Parse(timeFormat, str)

	return err
}
