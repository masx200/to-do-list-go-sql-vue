package main

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"math/rand"
)
import "gitee.com/masx200/to-do-list-go-sql-vue/backend/config"
import "gitee.com/masx200/to-do-list-go-sql-vue/backend/todoitem"
import "gitee.com/masx200/to-do-list-go-sql-vue/backend/database"
import "gitee.com/masx200/to-do-list-go-sql-vue/backend/operations"

func main() {
	var LoadConfig = config.LoadConfig
	type ToDoItem = todoitem.ToDoItem

	rand.Seed(time.Now().Unix())
	dsn := LoadConfig()
	db := database.ConnectDatabase(dsn, &ToDoItem{})

	fmt.Println("CreateItem", operations.CreateItem(db, &ToDoItem{Content: "hello world!" + strconv.FormatInt((rand.Int63n(math.MaxInt64)), 10), Finished: false}))
	fmt.Println("FindItems")
	fmt.Println(operations.FindItems(db, []ToDoItem{}, 30))
	fmt.Println("CreateItem", operations.CreateItem(db, &ToDoItem{Content: "hello world!" + strconv.FormatInt((rand.Int63n(math.MaxInt64)), 10), Finished: false}))
	fmt.Println("UpdateItem", operations.UpdateItem(db, &ToDoItem{Content: "changed!" + strconv.FormatInt((rand.Int63n(math.MaxInt64)), 10), Finished: true}, 2))
	fmt.Println("FindItems")
	fmt.Println(operations.FindItems(db, []ToDoItem{}, 30))
	fmt.Println("DeleteItem", operations.DeleteItem(db, &ToDoItem{}, 1))
}
