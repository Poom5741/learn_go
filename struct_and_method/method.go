package main

import "fmt"

type Book struct {
	Name   string
	Author string
}

func (book Book) String() string {
	return book.Name + " by " + book.Author
}
func main() {

	book := Book{
		Name:   "Harry Potter",
		Author: "J.K. Rowling",
	}
	fmt.Println(book.String())
}

func (book *Book) SetName(name string) string {
	book.Name = name
	return book.Name
}
