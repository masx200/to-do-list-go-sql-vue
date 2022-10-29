package database

import "gorm.io/gorm"

func CloneDB(db *gorm.DB) *gorm.DB {
	return db.Session(&gorm.Session{})
}
