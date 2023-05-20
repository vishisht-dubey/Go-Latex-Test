package section

type WorkExperience struct {
	Organisation string     `json:"organisation"`
	Role         string     `json:"role"`
	TimePeriod   TimePeriod `json:"timePeriod"`
	Links        []Link     `json:"links"`
	Details      []string   `json:"details"`
}
