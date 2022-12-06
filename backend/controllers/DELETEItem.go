package controllers

import (
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DELETEItem[T any](r *gin.Engine, GetDB func() (*gorm.DB, error), prefix string, model *T) {
	r.DELETE(prefix, func(c *gin.Context) {
		var db, err = GetDB()
		if err != nil {
			c.String(500, err.Error())
			return
		}
		var inputs []map[string]any
		err = c.ShouldBindJSON(&inputs)
		if err != nil {
			c.String(400, err.Error())
			return
		}
		ids := []uint{}
		for _, input := range inputs {
			qsid, o := input["id"]
			if !o {
				c.String(400, "expect id but not found")
				return
			}
			id, ok := qsid.(float64)

			if !ok {
				c.String(400, err.Error())
				return
			}
			ids = append(ids, uint(id))
		}
		err = database.DeleteItems(db, model, ids)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		var results = []map[string]interface{}{}
		for _, input := range ids {

			results = append(results, map[string]any{"id": input})
		}
		c.JSON(200, results)

	})
}
