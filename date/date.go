package date

import (
	"fmt"
	"strings"
	"time"
)

const ctLayout = "2006-01-02"

// Date is a time.Time representation with specific json marshaling capabilities to handle
// input dates considering of a string input "year-month-day"
type Date time.Time

// UnmarshalJSON implements the json.Unmarshaler interface. The date is expected to be in the "2006-01-02" format.
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(ctLayout, s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

// MarshalJSON implements the json.Marshaler interface. The date is a quoted string in in the "2006-01-02" format.
func (d Date) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	if t.UnixNano() == (time.Time{}).UnixNano() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, t.Format(ctLayout))), nil
}

func (d Date) Format(s string) string {
	t := time.Time(d)
	return t.Format(s)
}
