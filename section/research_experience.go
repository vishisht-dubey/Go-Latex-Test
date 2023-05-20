package section

import (
	"bytes"
	"text/template"
)

type ResearchExperience struct {
	ProjectTitle   string            `json:"projectTitle"`
	Institute      string            `json:"institute"`
	ProjectAdvisor string            `json:"projectAdvisor"`
	TimePeriod     TimePeriod        `json:"timePeriod"`
	Details        []MarkdownSnippet `json:"details"`
}

var tplResearchExperience = template.Must(template.ParseFiles("section/tex/research_experience.tex"))

func PrepareResearchExperienceSection(researchExperience []ResearchExperience) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	err := tplResearchExperience.Execute(&buf, researchExperience)
	if err != nil {
		println(err.Error())
		return bytes.NewBufferString(""), err
	}

	println(buf.String())
	return &buf, nil
}
