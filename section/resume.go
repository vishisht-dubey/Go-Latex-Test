package section

type Resume struct {
	Contact                   Contact              `json:"contact"`
	Skills                    []Skills             `json:"skills"`
	Education                 []Education          `json:"education"`
	WorkExperience            []WorkExperience     `json:"workExperience"`
	ResearchExperience        []ResearchExperience `json:"researchExperience"`
	PersonalProjects          []PersonalProjects   `json:"personalProjects"`
	Achievements              []MarkdownSnippet    `json:"achievements"`
	ExtraCurriculars          []MarkdownSnippet    `json:"extraCurriculars"`
	PositionsOfResponsibility []POR                `json:"positionsOfResponsibility"`
}
