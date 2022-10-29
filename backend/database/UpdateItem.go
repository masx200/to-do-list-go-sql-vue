package database

import (
	"gorm.io/gorm"
)

func UpdateItem[T any](db *gorm.DB, model *T, item map[string]any, id uint) error {

	db = CloneDB(db)
	// defer CloseDB(db)
	result := db.Model(&model).Where("id = ?", id).Select("*").Omit("id", "created_at", "deleted_at").Updates(&item)

	return result.Error
}
