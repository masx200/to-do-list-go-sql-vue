package database

import (
	"gorm.io/gorm"
)

func UpdateItem[T any](createDB func() *gorm.DB, item *T, id uint) error {

	db := createDB()
	var model T
	result := db.Model(&model).Where("id = ?", id).Select("*").Omit("id", "created_at", "deleted_at").Updates(&item)

	return result.Error
}
