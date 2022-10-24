package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func UpsertItem[T any](createDB func() *gorm.DB, model *T, item map[string]any, id uint) error {

	db := createDB()

	result := db.Model(&model).Select("*").Omit("created_at", "deleted_at").Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&item).Update("deleted_at", nil)

	return result.Error
}
