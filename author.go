package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Author - type
type Author struct {
	Name  string `xml:"Name" json:"name"`
	Books []Book `xml:"Book" json:"books"`
}

// MarshalAuthorJSON - Book to JSON string
func (a *Author) MarshalAuthorJSON() string {
	str, err := json.Marshal(a)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(str)
}

// MarshalAuthorXML - Book to XML string
func (a *Author) MarshalAuthorXML() string {
	str, err := xml.MarshalIndent(a, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(str)
}

/* unMarshaling for JSON and XML */

// UnmarshalAuthorXML - xml string to book
func (a *Author) UnmarshalAuthorXML(str string) {
	xml.Unmarshal([]byte(str), a)
}

// UnmarshalAuthorJSON - JSON string to book
func (a *Author) UnmarshalAuthorJSON(str string) {
	json.Unmarshal([]byte(str), a)
}
