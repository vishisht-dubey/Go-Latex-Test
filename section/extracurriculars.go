package section

import (
	"bytes"
	"text/template"
)

var tplextraCurriculars = template.Must(template.ParseFiles("section/tex/extracurriculars.tex"))

func PrepareExtraCurricularsSection(extraCurriculars []MarkdownSnippet) (*bytes.Buffer, error) {
	return PrepareSection(extraCurriculars, tplextraCurriculars)
}
