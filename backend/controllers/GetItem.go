package controllers

import (
	// "fmt"

	"gorm.io/gorm"
)

func GetItem[T any](db *gorm.DB, item *T, id uint) (*T, error) {
	// fmt.Println("update")
	// fmt.Print("\n\n")

	result := db.First(&item, id)
	// fmt.Printf("%#v\n", item)
	// fmt.Printf("%#v\n", result)

	return item, result.Error
}
