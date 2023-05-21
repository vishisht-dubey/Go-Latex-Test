package section

import (
	"bytes"
	"text/template"
)

type WorkExperience struct {
	Organisation string            `json:"organisation"`
	Location     string            `json:"location"`
	Role         string            `json:"role"`
	TimePeriod   TimePeriod        `json:"timePeriod"`
	Links        []Link            `json:"links"`
	Details      []MarkdownSnippet `json:"details"`
}

var tplWorkExperience = template.Must(template.ParseFiles("section/tex/work_experience.tex"))

func PrepareWorkExperience(workexperience []WorkExperience) (*bytes.Buffer, error) {
	return PrepareSection(workexperience, tplWorkExperience)
}
