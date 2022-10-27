package database

import (
	"gorm.io/gorm"
)

func FindItems[T any](createDB func() *gorm.DB, limit int, page int, model *T, query map[string]any, order string, direction string) ([]map[string]any, error) {
	db := createDB()
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	if direction == "desc" {
		db = db.Order(order + " " + "DESC")
	} else {
		db = db.Order(order + " " + "ASC")
	}
	var items = []map[string]any{}
	result := db.Model(model).Omit("deleted_at").Where(query).Limit(limit).Offset(page * limit).Find(&items)

	return items, result.Error
}

func FindByIDs[T any](createDB func() *gorm.DB, model *T, ids []uint) ([]map[string]any, error) {
	db := createDB()
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	var items = make([]map[string]any, len(ids))
	result := db.Model(model).Omit("deleted_at").Find(&items, ids)

	return items, result.Error
}
