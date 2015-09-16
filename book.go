package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Book - define bood
type Book struct {
	Title string `xml:"Title" json:"title"`
	Year  int    `xml:"Year" json:"year"`
}

/* This is added to implement sort.Interface to sort books by year */

// ByYear implements sort.Interface for []Book base on Year field
type ByYear []Book

// Len implements sort.Interface Len
func (b ByYear) Len() int { return len(b) }

// Swap implements sort.Interface for Less
func (b ByYear) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

// Less implements sort.Interface for Less
func (b ByYear) Less(i, j int) bool { return b[i].Year < b[j].Year }

/* Marshaling for JSON and XML */

// MarshalBookJSON - Book to JSON string
func (b *Book) MarshalBookJSON() string {
	str, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(str)
}

// MarshalBookXML - Book to XML string
func (b *Book) MarshalBookXML() string {
	str, err := xml.MarshalIndent(b, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(str)
}

/* unMarshaling for JSON and XML */

// UnmarshalBookXML - xml string to book
func (b *Book) UnmarshalBookXML(str string) {
	xml.Unmarshal([]byte(str), b)
}

// UnmarshalBookJSON - JSON string to book
func (b *Book) UnmarshalBookJSON(str string) {
	json.Unmarshal([]byte(str), b)
}
