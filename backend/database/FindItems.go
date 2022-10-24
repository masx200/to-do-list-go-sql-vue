package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func FindItems[T any](db *gorm.DB, limit int, page int, model *T, query *T, order string, direction string) ([]map[string]any, error) {

	var items = []map[string]any{}
	result := db.Model(model).Omit("deleted_at").Where(query).Limit(limit).Offset(page * limit).Order(clause.OrderByColumn{Column: clause.Column{Name: order}, Desc: direction == "desc"}).Find(&items)

	return items, result.Error
}
