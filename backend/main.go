package main

import (
	"strconv"

	"gitee.com/masx200/to-do-list-go-sql-vue/backend/configs"
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/database"
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/router"
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/todoitem"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	var LoadConfig = configs.LoadConfig
	type ToDoItem = todoitem.ToDoItem
	var TodoRoute = router.TodoRoute
	config := LoadConfig()
	db := database.ConnectDatabase(config.Dsn, &ToDoItem{})
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	TodoRoute(r, db)
	r.Run(":" + strconv.Itoa(config.Port))

}
