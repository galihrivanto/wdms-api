package wdmsapi

import (
	"strconv"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

type Time struct {
	time.Time

	// hours different from GMT
	Offset int
}

func (t *Time) String() string {
	return t.Time.String()
}

func (t *Time) location() *time.Location {
	return time.FixedZone("", t.Offset*3600)
}

func (t *Time) UnmarshalJSON(data []byte) error {
	str, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	// append second fraction if not exists
	if len(str) < len(timeFormat) {
		str += ":00"
	}

	t.Time, err = time.Parse(timeFormat, str)

	return err
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(t.Time.In(t.location()).Format(timeFormat))), nil
}

func (t Time) MarshalURLQuery() string {
	return t.Time.In(t.location()).Format(timeFormat)
}
