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

		limit, err := strconv.Atoi(qslimit)

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
	r.POST("/todoitem", func(c *gin.Context) {
		var item ToDoItem
		err := c.ShouldBindJSON(&item)
		if err != nil {
			c.String(400, err.Error())
			return
		}
		err = operations.CreateItem(db, &item)
		if err != nil {
			c.String(500, err.Error())
		} else {
			c.JSON(200, item)
		}

	})

	r.DELETE("/todoitem", func(c *gin.Context) {
		qsid := c.Query("id")
		if len(qsid) == 0 {
			c.String(400, "expect id not found")
			return
		}
		id, err := strconv.Atoi(qsid)

		if err != nil {
			c.String(400, err.Error())
			return
		}
		err = operations.DeleteItem(db, &ToDoItem{}, uint(id))
		if err != nil {
			c.String(500, err.Error())
		} else {
			c.JSON(200, gin.H{"id": id})
		}

	})
}
