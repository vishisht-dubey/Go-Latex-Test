package section

import (
	"bytes"
	"text/template"
)

type Skill struct {
	SkillType  string          `json:"skillType"`
	SkillItems MarkdownSnippet `json:"details"`
}

var tplSkills = template.Must(template.ParseFiles("section/tex/skills.tex"))

func PrepareSkillsSection(skills []Skill) (*bytes.Buffer, error) {
	return PrepareSection(skills, tplSkills)
}
