package inventory

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Item interface defines the methods for all item types
type Item interface {
	GetID() int
	GetName() string
	Details() string
}

// Inventory holds a slice of items
type Inventory struct {
	items map[int]Item
}

// NewInventory creates a new inventory
func NewInventory() *Inventory {
	return &Inventory{items: map[int]Item{}}
}

// AddItem adds a new item to the inventory
func (inv *Inventory) AddItem(item Item) error{
	// check if item already exists
	for _, i := range inv.items{
		if i.GetID() == item.GetID(){
			log.Warn("Attempted to add duplicate item")
			return errors.New("item already exists")
		}
	}

	inv.items[item.GetID()] = item
	log.Info("Item added successfully")
	return nil
}

// RemoveItem removes an item from the inventory
func (inv *Inventory) RemoveItem(id int) error {
	// Check if the item exists
	if _, exists := inv.items[id]; !exists {
		log.Warn("Attempted to remove non-existent item")
		return errors.New("item not found")
	}

	delete(inv.items, id)
	log.Info("Item removed successfully")
	return nil
}


// ListItems lists all items in the inventory
func (inv *Inventory) ListItems() {
	if len(inv.items) == 0 {
		log.Warn("Attempted to list empty inventory")
		return
	}

	fmt.Println("Items in inventory:")
	for _, item := range inv.items {
		fmt.Println(item.Details())
	}
}

