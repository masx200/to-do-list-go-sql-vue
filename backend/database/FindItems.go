package database

import (
	"gorm.io/gorm"
)

func FindItems[T any](db *gorm.DB, limit int, page int, model *T, query *T) ([]map[string]any, error) {

	var items = []map[string]any{}
	result := db.Model(model).Omit("deleted_at").Where(query).Limit(limit).Offset(page * limit).Find(&items)

	return items, result.Error
}
