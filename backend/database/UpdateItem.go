package database

import (
	// "fmt"

	"gorm.io/gorm"
)

func UpdateItem[T any](db *gorm.DB, item *T, id uint) error {
	// fmt.Println("update")
	// fmt.Print("\n\n")
	var model T
	result := db.Model(&model).Where("id = ?", id).Select("*").Omit("id", "created_at", "deleted_at").Updates(&item)
	// fmt.Printf("%#v\n", item)
	// fmt.Printf("%#v\n", result)

	return result.Error
}