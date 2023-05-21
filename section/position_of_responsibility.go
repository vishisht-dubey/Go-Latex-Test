package section

import (
	"bytes"
	"text/template"
)

type POR struct {
	Position         string            `json:"position"`
	Responsibilities []MarkdownSnippet `json:"responsibilities"`
	TimePeriod       TimePeriod        `json:"timePeriod"`
}

var tplPORs = template.Must(template.ParseFiles("section/tex/pors.tex"))

func PreparePORsSection(pors []POR) (*bytes.Buffer, error) {
	return PrepareSection(pors, tplPORs)
}
