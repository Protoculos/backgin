package domain

import "encoding/xml"

type Person struct {
	XMLName   xml.Name `json:"-" xml:"person,omitempty"`
	FirstName string   `json:"first_name,omitempty" xml:"first_name,omitempty"`
	LastName  string   `json:"last_name,omitempty" xml:"last_name,omitempty"`
}
