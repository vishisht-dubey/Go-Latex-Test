package section

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"
	"time"
)

type Date time.Time

type Education struct {
	StartDate   Date   `json:"start_date"`
	EndDate     Date   `json:"end_date"`
	Institution string `json:"institution"`
	Location    string `json:"Location"`
	Course      string `json:"course"`
	Degree      string `json:"degree"`
}

// Implement Marshaler and Unmarshaler interface
func (j *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = Date(t)
	return nil
}

func (j Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

var tpl = template.Must(template.ParseFiles("section/tex/education.tex"))

func PrepareEducationSection(e Education) (string, error) {
	var buf bytes.Buffer
	err := tpl.Execute(&buf, nil)
	if err != nil {
		// Handle the error if any and return
		fmt.Println(err.Error())
		return "", err
	}
	return buf.String(), nil
}
