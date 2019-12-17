package wdmsapi

import "strconv"

// Number represent number-string which used by some fields
// in WDMS api response, eg: "1", "20" for count
type Number struct {
	Value int
}

func (n Number) String() string {
	return strconv.Itoa(n.Value)
}

func (n Number) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(strconv.Itoa(n.Value))), nil
}

func (n *Number) UnmarshalJSON(data []byte) error {
	str, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	n.Value, err = strconv.Atoi(str)

	return err
}
