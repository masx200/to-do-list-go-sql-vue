package database

import (
	"errors"

	"gorm.io/gorm"
)

func GetItem[T any](createDB func() *gorm.DB, model *T, id uint) (map[string]interface{}, error) {
	db := createDB()

	users := []map[string]interface{}{}
	result := db.Limit(1).Model(&model).Omit("deleted_at").Where("id = ?", id).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(users) > 0 {
		return users[0], nil
	}
	return nil, errors.New("record not found")
}
