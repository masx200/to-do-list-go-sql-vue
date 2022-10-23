package controllers

import (
	"strconv"

	"gitee.com/masx200/to-do-list-go-sql-vue/backend/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DELETEItem[T any](r *gin.Engine, db *gorm.DB, prefix string, model *T) {
	r.DELETE(prefix, func(c *gin.Context) {
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
		res, err := database.GetItem(db, item, uint(id))
		/* 保持接口的幂等性 */
		if err != nil {
			c.JSON(200, []gin.H{{
				"id": id,
			}})
			return
		}
		err = database.DeleteItem(db, new(T), uint(id))
		if err != nil {
			c.String(500, err.Error())
		} else {
			c.JSON(200, []map[string]interface{}{res})
		}
		// return
	})
}
