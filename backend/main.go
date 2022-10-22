package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ToDoItem struct {
	gorm.Model
	Content string

	Finished bool
}

func main() {
	config := map[string]string{}
	x, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(x, &config)
	if err != nil {
		panic(err)
	}
	dsn, ok := config["dsn"]
	if !ok {
		panic(errors.New("config dsn not found"))
	}
	println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", db)
	err = db.AutoMigrate(&ToDoItem{})
	if err != nil {
		panic(err)
	}
	db.Create(&ToDoItem{Content: "hello world!", Finished: false})

}
