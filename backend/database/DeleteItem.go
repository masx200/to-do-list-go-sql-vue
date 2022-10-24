package database

import (
	"gorm.io/gorm"
)

func DeleteItem[T any](createDB func() *gorm.DB, model *T, id uint) error {
	db := createDB()

	return db.Model(&model).Delete(model, id).Error
}
