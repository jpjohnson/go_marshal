package main

import (
	"fmt"
	"testing"
)

// test new Book
func TestBook_Type(t *testing.T) {
	// Book 1 - type
	b1 := new(Book)
	b1.Title = "Adventures of Huckleberry Finn"
	b1.Year = 1884
	if b1.Year != 1884 {
		t.Fail()
	}
}

// test xml to Book
func TestBook_XML(t *testing.T) {
	// Book  - xml
	b := new(Book)
	b.UnmarshalBookXML(`<Book>
    <Title>The Prince and the Pauper</Title>
    <Year>1881</Year>
  </Book>`)
	if b.Year != 1881 {
		t.Fail()
	}
}

// test json to Book
func TestBook_JSON(t *testing.T) {
	// Book - JSON
	b := new(Book)
	b.UnmarshalBookJSON(`{"title":"The Adventures of Tom Sawyer","year":1876}`)
	if b.Year != 1876 {
		t.Fail()
	}
}

// example of book to xml
func ExampleBook_XML() {
	b := new(Book)
	b.Title = "Adventures of Huckleberry Finn"
	b.Year = 1884
	fmt.Println(b.MarshalBookXML())

	// Output:
	//<Book>
	//   <Title>Adventures of Huckleberry Finn</Title>
	//   <Year>1884</Year>
	//</Book>
}

// example of book to json
func ExampleBook_JSON() {
	b := new(Book)
	b.Title = "Adventures of Huckleberry Finn"
	b.Year = 1884
	fmt.Println(b.MarshalBookJSON())

	// Output:
	//{"title":"Adventures of Huckleberry Finn","year":1884}
}
