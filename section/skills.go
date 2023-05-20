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

func PrepareSkillsSection(skills []Skill) (string, error) {
	var buf bytes.Buffer
	err := tplSkills.Execute(&buf, skills)
	if err != nil {
		println(err.Error())
		return "", err
	}
	return buf.String(), nil
}
