package section

type PersonalProject struct {
	ProjectName string            `json:"projectName"`
	TimePeriod  TimePeriod        `json:"timePeriod"`
	Links       []Link            `json:"links"`
	Details     []MarkdownSnippet `json:"details"`
}
