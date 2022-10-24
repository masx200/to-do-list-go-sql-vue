package database

import (
	"gorm.io/gorm"
)

func GetItem[T any](createDB func() *gorm.DB, model *T, id uint) (map[string]interface{}, error) {
	db := createDB()

	res := map[string]interface{}{}
	result := db.Model(&model).Omit("deleted_at").First(res, id)

	return res, result.Error
}
