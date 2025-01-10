package main

import (
	"fmt"
	"inventory-management/inventory"
	"inventory-management/items"
)

// runPanic runs a function that panics
func runPanic(err error) {
	defer handlePanic()
	panic(err)
}

// // AddItem adds a new item to the inventory
func handlePanic(){
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic")
	}
}

func main() {
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
	}

	inv.ListItems()

}