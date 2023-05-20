package section

import (
	"bytes"
	"text/template"
)

type WorkExperience struct {
	Organisation string            `json:"organisation"`
	Role         string            `json:"role"`
	TimePeriod   TimePeriod        `json:"timePeriod"`
	Links        []Link            `json:"links"`
	Details      []MarkdownSnippet `json:"details"`
}

var tplWorkExperience = template.Must(template.ParseFiles("section/tex/work_experience.tex"))

func PrepareWorkExperience(workexperience []WorkExperience) (string, error) {
	var buf bytes.Buffer
	err :=  tplWorkExperience.Execute(&buf, workexperience)
	if err != nil {
		println(err.Error())
		return "", err
	}
	return buf.String(), nil
}
