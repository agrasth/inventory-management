package inventory

import "inventory-management/constants"

// Category is a custom type for item categories
type Category int32

// Convert Category to Category Name
func (c Category) ToString() string {
	switch int32(c) {
	case constants.BookCategory:
		return constants.BookCategoryName
	case constants.ElectronicCategory:
		return constants.ElectronicCategoryName
	default:
		return "Unknown"
	}
}

// Convert Category Name to Category
func StringToCategory(name string) Category {
	switch name {
	case constants.BookCategoryName:
		return Category(constants.BookCategory)
	case constants.ElectronicCategoryName:
		return Category(constants.ElectronicCategory)
	default:
		return -1 // Unknown category
	}
}
