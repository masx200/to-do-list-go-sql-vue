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

func UpdateItem[T any](db *gorm.DB, item *T, id uint) error {
	// fmt.Println("update")
	// fmt.Print("\n\n")
	var model T
	result := db.Model(&model).Where("id = ?", id).Select("*").Omit("id", "created_at", "deleted_at").Updates(&item)
	// fmt.Printf("%#v\n", item)
	// fmt.Printf("%#v\n", result)

	return result.Error
}

func CreateItem[T any](db *gorm.DB, model *T, item map[string]any) error {
	// fmt.Println("create")
	// fmt.Print("\n\n")
	delete(item, "id")
	result := db.Model(model).Create(&item)
	// fmt.Printf("%#v\n", item)
	// fmt.Printf("%#v\n", result)
	return result.Error
}
