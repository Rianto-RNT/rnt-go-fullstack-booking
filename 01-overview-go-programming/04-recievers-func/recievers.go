package main

import "fmt"

type Book struct {
	Title       string
	ReleaseYear int
}

func(b *Book) printTitle() string {
	return b.Title
}

func main() {
	var books Book
	books.Title = "Wonderful Island"

	newBook := Book{
		Title: "Poor Country",
	} 

	fmt.Println("Book title is", books.Title)
	fmt.Println("The new book title is", newBook.Title)
}