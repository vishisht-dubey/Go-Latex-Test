package section

type ResearchExperience struct {
	ProjectTitle   string     `json:"projectTitle"`
	Institute      string     `json:"institute"`
	Location       string     `json:"location"`
	ProjectAdvisor string     `json:"projectAdvisor"`
	TimePeriod     TimePeriod `json:"timePeriod"`
	Details        []string   `json:"details"`
}
