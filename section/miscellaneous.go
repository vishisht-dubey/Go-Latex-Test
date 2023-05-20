package section

import (
	"encoding/json"
	"strings"
	"time"
)

type Link struct {
	LinkURL  string `json:"linkURL"`
	LinkText string `json:"linkText"`
}

type TimePeriod struct {
	StartDate Date `json:"startDate"`
	EndDate   Date `json:"endDate"`
	IsOngoing bool `json:"isOngoing"`
}

type Date struct {
	raw_time    time.Time
	String_time string
}

// Implement Marshaler and Unmarshaler interface
func (j *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		println(err)
		return err
	}
	str := t.Format("%Y-%M-%D")
	*j = Date{
		t,
		str,
	}
	return nil
}

func (j Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j.raw_time))
}
