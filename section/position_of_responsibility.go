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

func PreparePORsSection(pors []POR) (string, error) {
	var buf bytes.Buffer
	err := tplPORs.Execute(&buf, pors)
	if err != nil {
		println(err.Error())
		return "", err
	}

	println(buf.String())
	return buf.String(), nil
}
