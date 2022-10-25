package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func UpsertItem[T any](createDB func() *gorm.DB, model *T, item map[string]any, id uint) error {
	var obj T
	empty := StructToMap(&obj)
	db := createDB()
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	for k, v := range item {
		empty[k] = v
	}
	delete(empty, "updated_at")
	result := db.Model(&model).Select("*").Omit("created_at", " updated_at").Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&empty).Where("id = ?", id).Unscoped().Update("deleted_at", nil)

	return result.Error
}
