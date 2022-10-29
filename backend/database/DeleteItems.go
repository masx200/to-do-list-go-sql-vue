package database

import (
	"gorm.io/gorm"
)

func DeleteItems[T any](db *gorm.DB, model *T, ids []uint) error {
	db = CloneDB(db)
	// defer CloseDB(db)
	return db.Model(&model).Delete(model, ids).Error
}

// func CloseDB(db *gorm.DB) error {
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return sqlDB.Close()
// }
