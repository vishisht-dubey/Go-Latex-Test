package section

import (
	"bytes"
	"text/template"
)

var tplAchievements = template.Must(template.ParseFiles("section/tex/achievements.tex"))

func PrepareAchievementsSection(achievements []MarkdownSnippet) (*bytes.Buffer, error) {
	return PrepareSection(achievements, tplAchievements)
}
