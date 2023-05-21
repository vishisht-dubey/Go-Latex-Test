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
	return PrepareSection(researchExperience, tplResearchExperience)
}
