package section

type Skills struct {
	SkillType  string            `json:"skillType"`
	SkillItems MarkdownSnippet `json:"details"`
}
