package controllers

import (
	// "fmt"
	"strconv"

	"gitee.com/masx200/to-do-list-go-sql-vue/backend/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GETItems[T any](r *gin.Engine, db *gorm.DB, prefix string, model *T) {
	r.GET(prefix, func(c *gin.Context) {

		qsid := c.Query("id")

		if len(qsid) != 0 {
			c.Abort()
			id, err := strconv.Atoi(qsid)

			if err != nil {
				c.String(400, err.Error())
				return
			}

			res, err := database.GetItem(db, model, uint(id))
			if err != nil {
				c.JSON(200, []map[string]interface{}{})
			} else {
				c.JSON(200, []map[string]interface{}{res})
			}
			return
		} else {
			c.Next()
		}

	},
		func(c *gin.Context) {
			qsid := c.Query("id")

			if len(qsid) != 0 {
				return
			}
			c.Abort()
			qslimit := c.DefaultQuery("limit", "50")

			limit, err := strconv.Atoi(qslimit)
			if err != nil {
				c.String(400, err.Error())
				return
			}
			qspage := c.DefaultQuery("page", "0")

			page, err := strconv.Atoi(qspage)

			if err != nil {
				c.String(400, err.Error())
				return
			}

			var query T

			err = c.ShouldBindQuery(&query)
			if err != nil {
				c.String(400, err.Error())
				return
			}
			// fmt.Printf("%#v\n", query)
			tdi, err := database.FindItems(db, limit, page, model, &query)
			if err != nil {
				c.String(500, err.Error())
			} else {
				c.JSON(200, tdi)
			}
			// return
		})
}
