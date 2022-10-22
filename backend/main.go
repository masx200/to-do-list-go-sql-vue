package main

import (
	"strconv"

	"gitee.com/masx200/to-do-list-go-sql-vue/backend/configs"
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/database"
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/operations"
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/todoitem"
	"github.com/gin-gonic/gin"
)

func main() {
	var LoadConfig = configs.LoadConfig
	type ToDoItem = todoitem.ToDoItem

	config := LoadConfig()
	db := database.ConnectDatabase(config.Dsn, &ToDoItem{})
	r := gin.Default()
	r.GET("/todoitem", func(c *gin.Context) {
		qslimit := c.DefaultQuery("limit", "30")
		var err error
		limit := 30

		limit, err = strconv.Atoi(qslimit)

		if err != nil {
			c.String(400, err.Error())
			return
		}
		tdi, err := operations.FindItems(db, []ToDoItem{}, limit)
		if err != nil {
			c.String(500, err.Error())
		} else {
			c.JSON(200, tdi)
		}

	})
	r.Run(":" + strconv.Itoa(config.Port))

}
