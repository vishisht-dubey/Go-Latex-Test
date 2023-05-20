package section

import (
	"Go-Latex-Test/shellescape"
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
	layout := "2006-01-02"
	t, err := time.Parse(layout, s)
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

type StringNotSafeError struct{}

func (m *StringNotSafeError) Error() string {
	return "given string is not safe to process"
}

func (m *MarkdownSnippet) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	println(shellescape.Quote(s))
	stdout, err := exec.Command("echo", s, " | pandoc -f markdown -t latex").Output()
	if err != nil {
		println("not successful")
		println(err.Error())
		return err
	}

	println("successful : ", stdout)

	*m = MarkdownSnippet{
		s,
		string(stdout),
	}

	return nil
}

func (m MarkdownSnippet) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.raw_string)
}
