package database

import "gorm.io/driver/mysql"
import "gorm.io/gorm"
import "fmt"

func ConnectDatabase(dsn string, model any) *gorm.DB {
	fmt.Println("connect")
	fmt.Print("\n\n")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", db)
	db = db.Debug()
	err = db.AutoMigrate(model)
	if err != nil {
		panic(err)
	}
	return db
}