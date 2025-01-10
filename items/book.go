package items

import "fmt"

// Book struct represents a book
type Book struct {
	ID     int
	Name   string
	Author string
}

//	GetID returns the ID of the book
func (b *Book) GetID() int {
	return b.ID
}

// 	GetName returns the name of the book
func (b *Book) GetName() string {
	return b.Name
}

//	Details returns the details of the book
func (b *Book) Details() string {
	return fmt.Sprintf("Book [ID: %d, Name: %s, Author: %s]", b.ID, b.Name, b.Author)
}