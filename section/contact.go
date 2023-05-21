package section

import (
	"bytes"
	"text/template"
)

var tplContact = template.Must(template.ParseFiles("section/tex/contact.tex"))

func PrepareContact(contact Contact) (string, error) {
	var buf bytes.Buffer
	err := tplContact.Execute(&buf, contact)
	if err != nil {
		println(err.Error())
		return "", err
	}
	println(buf.String())
	return buf.String(), nil
}