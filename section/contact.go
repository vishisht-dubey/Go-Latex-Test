package section

import (
	"bytes"
	"text/template"
)

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

var tplContact = template.Must(template.ParseFiles("section/tex/contact.tex"))

func PrepareContactSection(contact Contact) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	err := tplContact.Execute(&buf, contact)
	if err != nil {
		println(err.Error())
		return bytes.NewBufferString(""), err
	}
	println(buf.String())
	return &buf, nil
}
