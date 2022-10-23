package controllers

import (
	"encoding/json"
	"strconv"

	"gitee.com/masx200/to-do-list-go-sql-vue/backend/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func POSTItem[T any](r *gin.Engine, db *gorm.DB, prefix string, model *T) {
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
		id, err := database.CreateItem(db, model, &item)
		if err != nil {
			c.String(500, err.Error())
			return
		} else {

			res, err := database.GetItem(db, model, uint(id))
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

func PUTItem[T any](r *gin.Engine, db *gorm.DB, prefix string, model *T) {
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

		err = database.UpdateItem(db, &item, uint(id))
		if err != nil {
			c.String(500, err.Error())
			return
		}
		res, err := database.GetItem(db, item, uint(id))
		if err != nil {
			c.String(500, err.Error())
		} else {
			c.JSON(200, []map[string]any{res})
		}
		// return
	})
}
