package operations

import (
	"fmt"

	"gorm.io/gorm"
)

func DeleteItem[T any](db *gorm.DB, model *T, id uint) *gorm.DB {
	fmt.Println("delete")
	fmt.Print("\n\n")
	return db.Delete(model, id)
}

func UpdateItem[T any](db *gorm.DB, item *T, id uint) *gorm.DB {
	fmt.Println("update")
	fmt.Print("\n\n")
	var model T
	result := db.Model(&model).Where("id = ?", id).Select("*").Omit("id", "created_at", "deleted_at").Updates(&item)
	fmt.Printf("%#v\n", item)
	fmt.Printf("%#v\n", result)
	return result
}

func FindItems[T any](db *gorm.DB, items []T) ([]T, *gorm.DB) {
	fmt.Println("find")
	fmt.Print("\n\n")

	result := db.Limit(50).Find(&items)
	fmt.Printf("%#v\n", items)
	fmt.Printf("%#v\n", result)
	return items, result
}

func CreateItem[T any](db *gorm.DB, item *T) *gorm.DB {
	fmt.Println("create")
	fmt.Print("\n\n")

	result := db.Create(&item)
	fmt.Printf("%#v\n", item)
	fmt.Printf("%#v\n", result)
	return result
}
