package database

import (
	"gorm.io/gorm"
)

func CreateItem[T any](createDB func() *gorm.DB, model *T, item *T) (uint, error) {
	db := createDB()
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	var id uint
	result := db.Model(model).Select("*").Omit("id", "deleted_at").Create(item)

	id = JSONGetID(item)
	return id, result.Error
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
