package database

import (
	"gorm.io/gorm"
)

func DeleteItems[T any](db *gorm.DB, model *T, ids []uint) error {
	db = CloneDB(db)

	return db.Model(&model).Delete(model, ids).Error
}
