package database

import (
	// "fmt"

	"gorm.io/gorm"
)

func DeleteItem[T any](db *gorm.DB, model *T, id uint) error {
	// fmt.Println("delete")
	// fmt.Print("\n\n")
	return db.Model(&model).Delete(model, id).Error
}
