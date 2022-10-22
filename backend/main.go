package main

import (
	// "fmt"
	// "math"
	"strconv"
	"time"

	"math/rand"

	"gitee.com/masx200/to-do-list-go-sql-vue/backend/configs"
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/database"
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/operations"
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/todoitem"
	"github.com/gin-gonic/gin"
)

func main() {
	var LoadConfig = configs.LoadConfig
	type ToDoItem = todoitem.ToDoItem

	rand.Seed(time.Now().Unix())
	config := LoadConfig()
	db := database.ConnectDatabase(config.Dsn, &ToDoItem{})
	r := gin.Default()
	r.Run(":" + strconv.Itoa(config.Port))
	// fmt.Println("CreateItem", operations.CreateItem(db, &ToDoItem{Content: "hello world!" + strconv.FormatInt((rand.Int63n(math.MaxInt64)), 10), Finished: false}))
	// fmt.Println("FindItems")
	// fmt.Println(operations.FindItems(db, []ToDoItem{}, 30))
	// fmt.Println("CreateItem", operations.CreateItem(db, &ToDoItem{Content: "hello world!" + strconv.FormatInt((rand.Int63n(math.MaxInt64)), 10), Finished: false}))
	// fmt.Println("UpdateItem", operations.UpdateItem(db, &ToDoItem{Content: "changed!" + strconv.FormatInt((rand.Int63n(math.MaxInt64)), 10), Finished: true}, 2))
	// fmt.Println("FindItems")
	// fmt.Println(operations.FindItems(db, []ToDoItem{}, 30))
	// fmt.Println("DeleteItem", operations.DeleteItem(db, &ToDoItem{}, 1))
}
