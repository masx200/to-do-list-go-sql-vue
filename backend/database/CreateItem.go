package database

import (
	"github.com/akrennmair/slice"
	"gorm.io/gorm"
)

func CreateItems[T any](db *gorm.DB, model *T, items []*T) ([]uint, error) {
	db = CloneDB(db)
	// defer CloseDB(db)
	var ids []uint
	result := db.Model(model).Select("*").Omit("id", "deleted_at").Create(items)

	ids = slice.Map(items, func(o *T) uint { return JSONGetID(o) }) // JSONGetID(items)
	return ids, result.Error
}
func JSONGetID[T any](obj *T) uint {

	var id uint = 0
	m := StructToMap(obj)

	id6, ok := (m["id"]).(float64)
	if !ok {
		return id
	}
	return uint(id6)
}
