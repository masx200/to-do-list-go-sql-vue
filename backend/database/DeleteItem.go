package database

import (
	"gorm.io/gorm"
)

func DeleteItem[T any](createDB func() *gorm.DB, model *T, id uint) error {
	db := createDB()
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	return db.Model(&model).Delete(model, id).Error
}
