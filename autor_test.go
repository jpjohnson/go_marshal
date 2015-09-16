package main

import (
	"fmt"
	"testing"
)

// test author type
func TestAuthor(t *testing.T) {
	a := new(Author)
	a.Name = "Mark Twain"
	if a.Name != "Mark Twain" {
		t.Fail()
	}

	if len(a.Books) > 0 {
		t.Fail()
	}
}

// test author from json
func TestAuthor_JSON(t *testing.T) {
	a := new(Author)
	a.UnmarshalAuthorJSON(`{"name":"Mark Twain","books":[{"title":"Adventures of Huckleberry Finn","year":1884}]}`)
	if a.Books[0].Year != 1884 {
		t.Fail()
	}
}

// test author from xml
func TestAuthor_XML(t *testing.T) {
	a := new(Author)
	a.UnmarshalAuthorXML(`<Author><Name>Mark Twain</Name><Book><Title>Adventures of Huckleberry Finn</Title><Year>1884</Year></Book></Author>`)
	if a.Books[0].Year != 1884 {
		t.Fail()
	}
}

// example author from json
func ExampleAuthor_JSON1() {
	a := new(Author)
	a.Name = "Mark Twain"

	fmt.Println(a.MarshalAuthorJSON())

	// Output:
	//{"name":"Mark Twain","books":null}
}

// example author from xml
func ExampleAuthor_XML1() {
	a := new(Author)
	a.Name = "Mark Twain"

	fmt.Println(a.MarshalAuthorXML())

	// Output:
	//<Author>
	//   <Name>Mark Twain</Name>
	//</Author>
}

// example author from json with boook
func ExampleAuthor_JSON2() {
	a := new(Author)
	a.Name = "Mark Twain"
	b := new(Book)
	b.Title = "Adventures of Huckleberry Finn"
	b.Year = 1884
	a.Books = append(a.Books, *b)
	fmt.Println(a.MarshalAuthorJSON())

	// Output:
	//{"name":"Mark Twain","books":[{"title":"Adventures of Huckleberry Finn","year":1884}]}
}

// example author from xml with book
func ExampleAuthor_XML2() {
	a := new(Author)
	a.Name = "Mark Twain"
	b := new(Book)
	b.Title = "Adventures of Huckleberry Finn"
	b.Year = 1884
	a.Books = append(a.Books, *b)
	fmt.Println(a.MarshalAuthorXML())

	// Output:
	//<Author>
	//   <Name>Mark Twain</Name>
	//   <Book>
	//     <Title>Adventures of Huckleberry Finn</Title>
	//     <Year>1884</Year>
	//   </Book>
	//</Author>
}
