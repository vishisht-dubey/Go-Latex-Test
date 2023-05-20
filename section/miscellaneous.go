package section

import (
	"Go-Latex-Test/utils"
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
	rawTime           time.Time
	FormattedToString string
	Year              int
	Month             string
}

type MarkdownSnippet struct {
	rawString        string
	FormattedToLatex string
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
	yearMonthDay := t.Format("%Y-%M-%D")
	year := t.Year()
	month := t.Month().String()

	*d = Date{
		t,
		yearMonthDay,
		year,
		month,
	}
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d.rawTime))
}

func (m *MarkdownSnippet) UnmarshalJSON(b []byte) error {
	s := utils.Quote(strings.Trim(string(b), "\""))

	stdout, err := exec.Command("bash", "-c", "echo "+s+" | pandoc -f markdown -t latex").Output()
	if err != nil {
		println(err.Error())
		return err
	}

	*m = MarkdownSnippet{
		s,
		strings.TrimRight(string(stdout), "\r\n"),
	}

	return nil
}

func (m MarkdownSnippet) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.rawString)
}
