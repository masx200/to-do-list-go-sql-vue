package database

import (
	"gorm.io/gorm"
)

func DeleteItems[T any](createDB func() *gorm.DB, model *T, ids []uint) error {
	db := createDB()
	defer CloseDB(db)
	return db.Model(&model).Delete(model, ids).Error
}
func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	return sqlDB.Close()
}
