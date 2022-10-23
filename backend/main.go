package main

import (
	"strconv"

	"gitee.com/masx200/to-do-list-go-sql-vue/backend/configs"
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/database"
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/models"
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/routers"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	var LoadConfig = configs.LoadConfig
	type ToDoItem = models.ToDoItem

	config := LoadConfig()
	db := database.ConnectDatabase(config.Dsn, &ToDoItem{}, "to_do_items", config.Debug)

	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	routers.TodoRoute[ToDoItem](r, db, "/todoitem")
	r.GET("/", func(c *gin.Context) {

		c.String(200, "index")
	})
	r.Run(":" + strconv.Itoa(config.Port))

}
