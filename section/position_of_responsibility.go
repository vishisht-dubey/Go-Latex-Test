package section

type POR struct {
	Position        string     `json:"position"`
	Responsibilites []string   `json:"string"`
	TimePeriod      TimePeriod `json:"timePeriod"`
}
