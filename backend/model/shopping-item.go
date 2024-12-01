package model

import "time"

// ShoppingItem представляет элемент в списке покупок
type ShoppingItem struct {
	ID             uint      `json:"id" gorm:"primaryKey"` // ID уникальный идентификатор элемента
	ShoppingListID uint      `json:"shopping_list_id"`     // ShoppingListID идентификатор связанного списка покупок
	Name           string    `json:"name"`                 // Name название элемента
	Amount         *int      `json:"amount,omitempty"`     // Amount количество данного элемента (опционально)
	Completed      bool      `json:"completed"`            // Completed состояние элемента
	CreatedAt      time.Time `json:"created_at"`           // CreatedAt время создания записи
	UpdatedAt      time.Time `json:"updated_at"`           // UpdatedAt время последнего обновления записи
}

type ApiCreateShoppingItem struct {
	Name      string `json:"name"`      // Name название элемента
	Amount    *int   `json:"amount"`    // Amount количество данного элемента (опционально)
	Completed bool   `json:"completed"` // Completed состояние элемента
}
