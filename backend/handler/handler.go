package handler

import (
	"gorm.io/gorm"
)

// Handler используется для обработки запросов и хранения зависимости от базы данных
type Handler struct {
	DB *gorm.DB
}
