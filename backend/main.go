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
	Content string `json:"content" gorm:"not null"`

	Finished bool `json:"finished"`
	ID       uint `gorm:"primarykey" json:"id"`
}

func main() {
	rand.Seed(time.Now().Unix())
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
	fmt.Println("connect")
	fmt.Print("\n\n")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", db)
	err = db.AutoMigrate(&ToDoItem{})
	if err != nil {
		panic(err)
	}
	fmt.Println("create")
	fmt.Print("\n\n")
	item1 := &ToDoItem{Content: "hello world!" + strconv.FormatInt((rand.Int63n(math.MaxInt64)), 10), Finished: false}
	result := db.Create(&item1)
	fmt.Printf("%#v\n", item1)
	fmt.Printf("%#v\n", result)
	fmt.Println("find")
	fmt.Print("\n\n")
	items := []ToDoItem{}
	result = db.Limit(50).Find(&items)
	fmt.Printf("%#v\n", items)
	fmt.Printf("%#v\n", result)
}
