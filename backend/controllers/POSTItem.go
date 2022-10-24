package controllers

import (
	"encoding/json"

	"gitee.com/masx200/to-do-list-go-sql-vue/backend/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func POSTItem[T any](r *gin.Engine, createDB func() *gorm.DB, prefix string, model *T) {
	r.POST(prefix, func(c *gin.Context) {
		var input map[string]any
		err := c.ShouldBindJSON(&input)
		if err != nil {
			c.String(400, err.Error())
			return
		}
		delete(input, "id")
		str, err := json.Marshal(input)
		if err != nil {
			c.String(400, err.Error())
			return
		}
		var item T

		err = json.Unmarshal(str, &item)
		if err != nil {
			c.String(400, err.Error())
			return
		}
		id, err := database.CreateItem(createDB, model, &item)
		if err != nil {
			c.String(500, err.Error())
			return
		} else {

			res, err := database.GetItem(createDB, model, uint(id))
			if err != nil {
				c.JSON(200, []map[string]interface{}{})
			} else {
				c.JSON(200, []map[string]interface{}{res})
			}
			return

		}
		// return
	})
}
