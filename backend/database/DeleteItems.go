package database

import (
	"gorm.io/gorm"
)

func DeleteItems[T any](createDB func() *gorm.DB, model *T, ids []uint) error {
	db := createDB()
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	return db.Model(&model).Delete(model, ids).Error
}
