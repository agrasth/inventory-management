package inventory

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Item defines a generic item with necessary fields
type Item[T any] struct {
	ID          int
	Name        string
	Category    Category
	CategoryName string
	Price       float32
	Quantity    int64
	Details     T
}

// Inventory holds a map of items with concurrency support
type Inventory struct {
	items map[int]*Item[any]
	mu    sync.RWMutex
}

// NewInventory creates a new inventory
func NewInventory() *Inventory {
	return &Inventory{items: make(map[int]*Item[any])}
}

// AddItem adds a new item to the inventory
func (inv *Inventory) AddItem(item *Item[any]) error {
	inv.mu.Lock()
	defer inv.mu.Unlock()

	if _, exists := inv.items[item.ID]; exists {
		log.Warn("Attempted to add duplicate item")
		return errors.New("item already exists")
	}
	item.CategoryName = item.Category.ToString()
	inv.items[item.ID] = item
	log.Info("Item added successfully")
	return nil
}

// RemoveItem removes an item from the inventory
func (inv *Inventory) RemoveItem(id int) error {
	inv.mu.Lock()
	defer inv.mu.Unlock()

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
	inv.mu.RLock()
	defer inv.mu.RUnlock()

	if len(inv.items) == 0 {
		log.Warn("Inventory is empty")
		return
	}

	fmt.Println("Items in inventory:")
	for _, item := range inv.items {
		fmt.Printf("ID: %d, Name: %s, Category: %s, Price: %.2f, Quantity: %d\n", item.ID, item.Name, item.CategoryName, item.Price, item.Quantity)
	}
}

// TopItemsByStock lists the top 5 items by stock quantity
func (inv *Inventory) TopItemsByStock() {
	inv.mu.RLock()
	defer inv.mu.RUnlock()

	var itemsSlice []*Item[any]
	for _, item := range inv.items {
		itemsSlice = append(itemsSlice, item)
	}

	sort.Slice(itemsSlice, func(i, j int) bool {
		return itemsSlice[i].Quantity > itemsSlice[j].Quantity
	})

	fmt.Println("Top 5 items by stock:")
	for i, item := range itemsSlice {
		if i >= 5 {
			break
		}
		fmt.Printf("ID: %d, Name: %s, Quantity: %d\n", item.ID, item.Name, item.Quantity)
	}
}
