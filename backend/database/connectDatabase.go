package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase[T any](dsn string, model *T, TableName string, debug bool) func() *gorm.DB {

	var createDB = func() *gorm.DB {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		db = db.Table(TableName)
		if debug {
			db = db.Debug()
		}

		db = db.Model(model)
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(100)
		return db
	}

	db := createDB()
	err := db.AutoMigrate(model)
	if err != nil {
		panic(err)
	}

	CloseDB(db)
	return createDB
}
