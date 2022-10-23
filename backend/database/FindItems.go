package database

import (
	// "fmt"

	"gorm.io/gorm"
)

func FindItems[T any](db *gorm.DB, limit int, page int, query *T) ([]map[string]any, error) {
	// fmt.Println("find")
	// fmt.Print("\n\n")
	var items = []map[string]any{}
	result := db.Model(query).Where(query).Limit(limit).Offset(page * limit).Find(&items)
	// fmt.Printf("%#v\n", items)
	// fmt.Printf("%#v\n", result)
	return items, result.Error
}
