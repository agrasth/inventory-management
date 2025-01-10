package main

import (
	"fmt"
	"github.com/agrasth/inventory-management/inventory"
	"github.com/agrasth/inventory-management/items"
)

// runPanic runs a function that panics
func runPanic(err error) {
	defer handlePanic()
	panic(err)
}

// handlePanic handles panics gracefully
func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func main() {
	// Create a new inventory
	inv := inventory.NewInventory()

	// Add items to the inventory
	err := inv.AddItem(&items.Book{ID: 1, Name: "Book 1", Author: "Author 1"})
	if err != nil {
		fmt.Println(err)
		runPanic(err)
	}

	err = inv.AddItem(&items.Electronic{ID: 2, Name: "Electronic 1", Brand: "Brand 1"})
	if err != nil {
		fmt.Println(err)
		runPanic(err)
	}

	// List items in the inventory
	inv.ListItems()

	// Remove an item from the inventory
	err = inv.RemoveItem(1)
	if err != nil {
		fmt.Println(err)
		runPanic(err)
	}

	// List remaining items
	inv.ListItems()
}
