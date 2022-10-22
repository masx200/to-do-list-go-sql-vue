package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"math/rand"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ToDoItem struct {
	gorm.Model
	Content string ` gorm:"not null"`

	Finished bool ` gorm:"not null"`
}

func main() {
	rand.Seed(time.Now().Unix())
	dsn := loadConfig()
	db := connectDatabase(dsn)

	fmt.Println("createItem", createItem(db, &ToDoItem{Content: "hello world!" + strconv.FormatInt((rand.Int63n(math.MaxInt64)), 10), Finished: false}))
	fmt.Println("findItems")
	fmt.Println(findItems(db))
	fmt.Println("createItem", createItem(db, &ToDoItem{Content: "hello world!" + strconv.FormatInt((rand.Int63n(math.MaxInt64)), 10), Finished: false}))
	fmt.Println("updateItem", updateItem(db, 2, &ToDoItem{Content: "changed!" + strconv.FormatInt((rand.Int63n(math.MaxInt64)), 10), Finished: true}))
	fmt.Println("findItems")
	fmt.Println(findItems(db))
	fmt.Println("deleteItem", deleteItem(db, 1))
}
func deleteItem(db *gorm.DB, id uint) *gorm.DB {
	fmt.Println("delete")
	fmt.Print("\n\n")
	return db.Delete(&ToDoItem{}, id)
}
func updateItem(db *gorm.DB, id uint, item1 *ToDoItem) *gorm.DB {
	fmt.Println("update")
	fmt.Print("\n\n")

	item := &ToDoItem{}
	item.ID = id
	result := db.Model(&item).Select("*").Omit("id", "created_at", "deleted_at").Updates(&item1)
	fmt.Printf("%#v\n", item1)
	fmt.Printf("%#v\n", result)
	return result
}

func findItems(db *gorm.DB) ([]ToDoItem, *gorm.DB) {
	fmt.Println("find")
	fmt.Print("\n\n")
	items := []ToDoItem{}
	result := db.Limit(50).Find(&items)
	fmt.Printf("%#v\n", items)
	fmt.Printf("%#v\n", result)
	return items, result
}

func createItem(db *gorm.DB, item1 *ToDoItem) *gorm.DB {
	fmt.Println("create")
	fmt.Print("\n\n")

	result := db.Create(&item1)
	fmt.Printf("%#v\n", item1)
	fmt.Printf("%#v\n", result)
	return result
}

func connectDatabase(dsn string) *gorm.DB {
	fmt.Println("connect")
	fmt.Print("\n\n")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", db)
	db = db.Debug()
	err = db.AutoMigrate(&ToDoItem{})
	if err != nil {
		panic(err)
	}
	return db
}

func loadConfig() string {
	fmt.Println("config")
	fmt.Print("\n\n")
	config := map[string]string{}
	text, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(text, &config)
	if err != nil {
		panic(err)
	}
	dsn, ok := config["dsn"]
	if !ok {
		panic(errors.New("config dsn not found"))
	}
	println(dsn)
	return dsn
}
