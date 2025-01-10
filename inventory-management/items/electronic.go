package items

import "fmt"

// Electronic struct represents an electronic item
type Electronic struct {
	ID    int
	Name  string
	Brand string
}

// GetID returns the ID of the electronic item
func (e *Electronic) GetID() int {
	return e.ID
}

// GetName returns the name of the electronic item
func (e *Electronic) GetName() string {
	return e.Name
}

// Details returns the details of the electronic item
func (e *Electronic) Details() string {
	return fmt.Sprintf("Book [ID: %d, Name: %s, Author: %s]", e.ID, e.Name, e.Brand)
}
