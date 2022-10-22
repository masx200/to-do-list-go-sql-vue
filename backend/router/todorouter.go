package router

import (
	"strconv"

	"gitee.com/masx200/to-do-list-go-sql-vue/backend/operations"
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/todoitem"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TodoRoute(r *gin.Engine, db *gorm.DB) {
	type ToDoItem = todoitem.ToDoItem
	r.GET("/todoitem", func(c *gin.Context) {
		qslimit := c.DefaultQuery("limit", "50")
		

		limit, err:= strconv.Atoi(qslimit)

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
}
