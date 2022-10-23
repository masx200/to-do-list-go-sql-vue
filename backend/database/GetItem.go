package database

import (
	// "fmt"

	"gorm.io/gorm"
)

func GetItem[T any](db *gorm.DB, model *T, id uint) (map[string]interface{}, error) {
	// fmt.Println("update")
	// fmt.Print("\n\n")
	res := map[string]interface{}{}
	result := db.Model(&model).Omit("deleted_at").First(res, id)
	// fmt.Printf("%#v\n", item)
	// fmt.Printf("%#v\n", result)

	return res, result.Error
}
