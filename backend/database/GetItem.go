package database

import (
	"errors"

	"gorm.io/gorm"
)

func GetItem[T any](db *gorm.DB, model *T, id uint) (map[string]interface{}, error) {
	db = CloneDB(db)
	// defer CloseDB(db)
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
