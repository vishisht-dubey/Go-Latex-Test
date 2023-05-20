package section

type PersonalProjects struct {
	ProjectName string     `json:"projectName"`
	TimePeriod  TimePeriod `json:"timePeriod"`
	Links       []Link     `json:"links"`
}
