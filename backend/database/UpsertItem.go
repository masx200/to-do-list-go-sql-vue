package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func UpsertItem[T any](db *gorm.DB, model *T, item map[string]any, id uint) error {
	db = CloneDB(db)
	var obj T
	empty := StructToMap(&obj)

	// defer CloseDB(db)
	for k, v := range item {
		empty[k] = v
	}
	delete(empty, "updated_at")
	result := db.Model(&model).Select("*").Omit("created_at", " updated_at").Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&empty).Where("id = ?", id).Unscoped().Update("deleted_at", nil)

	return result.Error
}
