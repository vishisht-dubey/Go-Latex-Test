package section

import (
	"bytes"
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

var tplEducation = template.Must(template.ParseFiles("section/tex/education.tex"))

func PrepareEducationSection(education []Education) (string, error) {
	var buf bytes.Buffer
	err := tplEducation.Execute(&buf, education)
	if err != nil {
		println(err.Error())
		return "", err
	}
	return buf.String(), nil
}
