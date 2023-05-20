package section

import (
	"encoding/json"
	"os/exec"
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
	raw_time       time.Time
	formatted_time string
}

type MarkdownSnippet struct {
	raw_string             string
	formatted_latex_string string
}

// Implement Marshaler and Unmarshaler interface
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		println(err)
		return err
	}
	str := t.Format("%Y-%M-%D")
	*d = Date{
		t,
		str,
	}
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d.raw_time))
}

func (m *MarkdownSnippet) UnmarshalJSON(b []byte) error {
	s := string(b)
	out, err := exec.Command("pandoc -f markdown -t latex <<< " + s).Output()
	if err != nil {
		println(err)
		return err
	}

	*m = MarkdownSnippet{
		s,
		string(out),
	}

	return nil
}

func (m MarkdownSnippet) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.raw_string)
}
