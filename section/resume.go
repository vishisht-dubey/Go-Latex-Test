package section

type Resume struct {
	Contact            Contact              `json:"contact"`
	Skills             []Skill              `json:"skills"`
	Education          []Education          `json:"education"`
	WorkExperience     []WorkExperience     `json:"workExperience"`
	ResearchExperience []ResearchExperience `json:"researchExperience"`
	PersonalProjects   []PersonalProject    `json:"personalProjects"`
	Achievements       []MarkdownSnippet    `json:"achievements"`
	ExtraCurriculars   []MarkdownSnippet    `json:"extraCurriculars"`
	PORs               []POR                `json:"positionsOfResponsibility"`
}
