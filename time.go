package nationbuilder

import (
	"strings"
	"time"
)

// The time format used by nationbuilder - usable in a call to format
const TimeFormat = "2006-01-02T15:04:05-07:00"

// A wrapper around time.Time to allow deserialising a string into a time object
type NationDate struct {
	Time *time.Time
}

// A nationbuilder formatted representation of the time
func (n *NationDate) String() string {
	if n.Time == nil {
		return ""
	}

	return n.Time.Format(TimeFormat)
}

// Implement json.Marshaller
func (n *NationDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + n.String() + `"`), nil
}

// Implement json.Unmarshaller
func (n *NationDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" {
		return nil
	}
	t, err := time.Parse(TimeFormat, s)
	if err != nil {
		return err
	}
	n.Time = &t
	return nil
}

// Create a new NationDate from a string representation of a date
// which follows nationbuilder's date format
func NewNationDate(date string) (*NationDate, error) {
	t, err := time.Parse(TimeFormat, date)
	if err != nil {
		return nil, err
	}
	return &NationDate{
		Time: &t,
	}, nil
}

// Shorthand function to return a nationdate wrapper around a time.Time object
func NewNationDateFromTime(t time.Time) *NationDate {
	return &NationDate{
		Time: &t,
	}
}
