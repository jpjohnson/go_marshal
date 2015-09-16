package main

import (
	"fmt"
	"sort"
)

func main() {

	// Book 1 - type
	b1 := new(Book)
	b1.Title = "Adventures of Huckleberry Finn"
	b1.Year = 1884
	// Book 2 - type
	b2 := new(Book)
	b2.Title = "The Adventures of Tom Sawyer"
	b2.Year = 1876
	// Book 3 - xml
	b3 := new(Book)
	b3.UnmarshalBookXML(`<Book>
    <Title>The Prince and the Pauper</Title>
    <Year>1881</Year>
  </Book>`)
	// Book 4 - JSON
	b4 := new(Book)
	b4.UnmarshalBookJSON(`{"title":"Roughing It","year":1872}`)

	// marshal json and xml
	bJSON := b1.MarshalBookJSON()
	bXML := b1.MarshalBookXML()
	fmt.Println(string(bJSON))
	fmt.Println(string(bXML))

	// author
	a := new(Author)
	a.Name = "Mark Twain"
	// add books
	a.Books = append(a.Books, *b1, *b2, *b3, *b4)
	// sort books by year, see book.go for sort.Interface implementation
	sort.Sort(ByYear(a.Books))

	// marshal json and xml
	aJSON := a.MarshalAuthorJSON()
	aXML := a.MarshalAuthorXML()
	fmt.Println(string(aJSON))
	fmt.Println(string(aXML))
}
