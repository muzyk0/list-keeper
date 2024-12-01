package store

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"list-keeper-backend/model"
)

func InitDatabase(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	migrateError := db.AutoMigrate(&model.ShoppingList{}, &model.ShoppingItem{})
	if migrateError != nil {
		panic(migrateError)
	}
	return db, nil
}
