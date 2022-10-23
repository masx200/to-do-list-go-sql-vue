package controllers

import (
	// "fmt"

	"gorm.io/gorm"
)

func DeleteItem[T any](db *gorm.DB, model *T, id uint) error {
	// fmt.Println("delete")
	// fmt.Print("\n\n")
	return db.Delete(model, id).Error
}

func UpdateItem[T any](db *gorm.DB, item *T, id uint) error {
	// fmt.Println("update")
	// fmt.Print("\n\n")
	var model T
	result := db.Model(&model).Where("id = ?", id).Select("*").Omit("id", "created_at", "deleted_at").Updates(&item)
	// fmt.Printf("%#v\n", item)
	// fmt.Printf("%#v\n", result)

	return result.Error
}

func FindItems[T any](db *gorm.DB, items []T, limit int, page int) ([]T, error) {
	// fmt.Println("find")
	// fmt.Print("\n\n")

	result := db.Limit(limit).Offset(page * limit).Find(&items)
	// fmt.Printf("%#v\n", items)
	// fmt.Printf("%#v\n", result)
	return items, result.Error
}

func CreateItem[T any](db *gorm.DB, item *T) error {
	// fmt.Println("create")
	// fmt.Print("\n\n")

	result := db.Create(&item)
	// fmt.Printf("%#v\n", item)
	// fmt.Printf("%#v\n", result)
	return result.Error
}
