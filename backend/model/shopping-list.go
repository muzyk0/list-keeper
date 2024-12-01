package model

import "time"

type ShoppingList struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `json:"name"`
	Completed bool           `json:"completed"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Items     []ShoppingItem `gorm:"foreignKey:ShoppingListID"`
}
