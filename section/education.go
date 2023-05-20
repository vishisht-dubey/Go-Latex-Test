package section

import (
	"bytes"
	"fmt"
	"text/template"
)

type Education struct {
	TimePeriod   TimePeriod        `json:"timePeriod"`
	Institution  string            `json:"institution"`
	Location     string            `json:"location"`
	Course       string            `json:"course"`
	Degree       string            `json:"degree"`
	MaximumMarks float32           `json:"maximumGPA"`
	MarksOrCGPA  float32           `json:"marksOrCGPA"`
	Details      []MarkdownSnippet `json:"details"`
}

var tpl = template.Must(template.ParseFiles("section/tex/education.tex"))

func PrepareEducationSection(e Education) (string, error) {
	var buf bytes.Buffer
	err := tpl.Execute(&buf, e)
	if err != nil {
		// Handle the error if any and return
		fmt.Println(err.Error())
		return "", err
	}
	return buf.String(), nil
}
