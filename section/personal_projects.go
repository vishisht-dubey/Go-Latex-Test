package section

import (
	"bytes"
	"text/template"
)

type PersonalProject struct {
	ProjectName string            `json:"projectName"`
	TimePeriod  TimePeriod        `json:"timePeriod"`
	Links       []Link            `json:"links"`
	Details     []MarkdownSnippet `json:"details"`
}

var tplPersonalProjects = template.Must(template.ParseFiles("section/tex/personal_projects.tex"))

func PreparePersonalProjectsSection(personalProjects []PersonalProject) (string, error) {
	var buf bytes.Buffer
	err := tplPersonalProjects.Execute(&buf, personalProjects)
	if err != nil {
		println(err.Error())
		return "", err
	}

	println(buf.String())
	return buf.String(), nil
}
