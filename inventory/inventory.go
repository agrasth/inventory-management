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
	items []Item
}

// NewInventory creates a new inventory
func NewInventory() *Inventory {
	return &Inventory{items: []Item{}}
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

	inv.items = append(inv.items, item)
	log.Info("Item added successfully")
	return nil
}

// RemoveItem removes an item from the inventory
func (inv *Inventory) RemoveItem(id int) error {
	for index, i := range inv.items {
		if i.GetID() == id {
			inv.items = append(inv.items[:index], inv.items[index+1:]...)
			log.Info("Item Removed successfully")
			return nil
		}
	}
	log.Warn("Attempted to remove non-existent item")
	return errors.New("item not found")
}

// ListItems lists all items in the inventory
func (inv* Inventory) ListItems() {
	if len(inv.items) == 0 {
		log.Warn("Attempted list empty inventory")
		return
	}

	fmt.Println("Items in inventory:")
	for _,i := range inv.items {
		fmt.Println(i.Details())
	}
}

