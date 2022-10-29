package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase[T any](dsn string, model *T, TableName string, debug bool) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = db.Table(TableName)
	if debug {
		db = db.Debug()
	}
	db = db.Model(model).Session(&gorm.Session{})

	// defer CloseDB(db)
	err = db.AutoMigrate(model)
	if err != nil {
		panic(err)
	}

	return db
}
