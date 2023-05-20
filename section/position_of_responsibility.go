package section

type POR struct {
	Position        string            `json:"position"`
	Responsibilites []MarkdownSnippet `json:"responsibilites"`
	TimePeriod      TimePeriod        `json:"timePeriod"`
}
