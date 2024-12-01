package model

import "time"

type ShoppingList struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Items     []ShoppingItem `json:"items" gorm:"foreignKey:ShoppingListID"`
}

type ApiCreateShoppingList struct {
	Name  string                  `json:"name"`
	Items []ApiCreateShoppingItem `json:"items"`
}
