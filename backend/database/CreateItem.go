package database

import (
	"encoding/json"

	"gorm.io/gorm"
)

func CreateItem[T any](createDB func() *gorm.DB, model *T, item *T) (uint, error) {
	db := createDB()

	var id uint
	result := db.Model(model).Select("*").Omit("id", "deleted_at").Create(item)

	id = JSONGetID(item)
	return id, result.Error
}
func JSONGetID(obj any) uint {
	str, err := json.Marshal(obj)

	var id uint = 0

	if err != nil {
		panic(err)
	}
	var m = map[string]any{}
	err = json.Unmarshal(str, &m)
	if err != nil {
		panic(err)
	}
	id6, ok := (m["id"]).(float64)
	if !ok {
		return id
	}
	return uint(id6)
}
