package database

import (
	"gorm.io/gorm"
)

func FindItems[T any](db *gorm.DB, limit int, page int, model *T, query map[string]any, order string, direction string) ([]map[string]any, error) {
	db = CloneDB(db)

	if direction == "desc" {
		db = db.Order(order + " " + "DESC")
	} else {
		db = db.Order(order + " " + "ASC")
	}
	var items = []map[string]any{}
	result := db.Model(model).Omit("deleted_at").Where(query).Limit(limit).Offset(page * limit).Find(&items)

	return items, result.Error
}

func FindByIDs[T any](db *gorm.DB, model *T, ids []uint) ([]map[string]any, error) {
	db = CloneDB(db)

	var items = make([]map[string]any, 0)
	result := db.Model(model).Omit("deleted_at").Find(&items, ids)

	return items, result.Error
}
