package controllers

import (
	"strconv"

	"gitee.com/masx200/to-do-list-go-sql-vue/backend/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PUTItem[T any](r *gin.Engine, createDB func() *gorm.DB, prefix string, model *T) {
	r.PUT(prefix, func(c *gin.Context) {
		qsid := c.Query("id")
		if len(qsid) == 0 {
			c.String(400, "expect id but not found")
			return
		}
		id, err := strconv.Atoi(qsid)

		if err != nil {
			c.String(400, err.Error())
			return
		}
		var item = new(T)
		err = c.ShouldBindJSON(&item)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		err = database.UpdateItem(createDB, &item, uint(id))
		if err != nil {
			c.String(500, err.Error())
			return
		}
		res, err := database.GetItem(createDB, item, uint(id))
		if err != nil {
			c.String(500, err.Error())
		} else {
			c.JSON(200, []map[string]any{res})
		}
		// return
	})
}
