package database

import (
	// "fmt"

	"encoding/json"
	// "fmt"

	"gorm.io/gorm"
)

func CreateItem[T any](db *gorm.DB, model *T, item *T) (uint, error) {
	// fmt.Println("create")
	// fmt.Print("\n\n")
	var id uint
	result := db.Model(model).Select("*").Omit("id", "created_at", "deleted_at").Create(item)
	// fmt.Printf("%#v\n", item)
	// fmt.Printf("%#v\n", result)
	id = JSONGetID(item)
	return id, result.Error
}
func JSONGetID(obj any) uint {
	str, err := json.Marshal(obj)

	var id uint = 0
	// fmt.Println(string(str), err)
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
