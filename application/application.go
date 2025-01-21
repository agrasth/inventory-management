package application

import (
	"github.com/agrasth/inventory-management/inventory"
)

func Run() {
	inv := inventory.NewInventory()

	// Add items
	_ = inv.AddItem(&inventory.Item[any]{ID: 1, Name: "Book 1", Category: inventory.Category(1), Price: 10.99, Quantity: 100})
	_ = inv.AddItem(&inventory.Item[any]{ID: 2, Name: "Electronic 1", Category: inventory.Category(2), Price: 299.99, Quantity: 50})

	// List items
	inv.ListItems()

	// Display top 5 items by stock
	inv.TopItemsByStock()

	// Remove an item
	_ = inv.RemoveItem(1)

	// List items again
	inv.ListItems()
}
