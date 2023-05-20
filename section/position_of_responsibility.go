package section

type POR struct {
	Position        string     `json:"position"`
	Responsibilites []string   `json:"responsibilites"`
	TimePeriod      TimePeriod `json:"timePeriod"`
}
