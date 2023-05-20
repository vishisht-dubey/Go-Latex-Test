package section

type ResearchExperience struct {
	ProjectTitle   string   `json:"projectTitle"`
	Institute      string   `json:"institute"`
	Location       string   `json:"location"`
	ProjectAdvisor string   `json:"projectAdvisor"`
	StartDate      string   `json:"startDate"`
	EndDate        string   `json:"endDate"`
	IsOngoing      string   `json:"isOngoing"`
	Details        []string `json:"details"`
}
