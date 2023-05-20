package section

import (
	"bytes"
	"text/template"
)

var tplAchievements = template.Must(template.ParseFiles("section/tex/achievements.tex"))

func PrepareAchievements(achievements []MarkdownSnippet) (string, error) {
	var buf bytes.Buffer
	err := tplAchievements.Execute(&buf, achievements)
	if err != nil {
		println(err.Error())
		return "", err
	}
	println(buf.String())
	return buf.String(), nil
}