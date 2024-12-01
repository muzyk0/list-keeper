package model

import "time"

// ShoppingItem представляет элемент в списке покупок
type ShoppingItem struct {
	ID             uint      `gorm:"primaryKey"`       // ID уникальный идентификатор элемента
	ShoppingListID uint      `json:"shopping_list_id"` // ShoppingListID идентификатор связанного списка покупок
	Name           string    `json:"name"`             // Name название элемента
	Amount         *int      `json:"amount,omitempty"` // Amount количество данного элемента (опционально)
	CreatedAt      time.Time `json:"created_at"`       // CreatedAt время создания записи
	UpdatedAt      time.Time `json:"updated_at"`       // UpdatedAt время последнего обновления записи
}
