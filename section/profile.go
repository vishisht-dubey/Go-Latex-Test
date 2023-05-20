package section

type Contact struct {
	Name         string         `json:"name"`
	Address      string         `json:"address"`
	Email        string         `json:"email"`
	PhoneNumber  string         `json:"phoneNumber"`
	ProfileLinks []ProfileLinks `json:"profileLinks"`
}

type ProfileLinks struct {
	IconName    string `json:"icon"`
	WebsiteName string `json:"websiteName"`
	Link        Link   `json:"link"`
}
