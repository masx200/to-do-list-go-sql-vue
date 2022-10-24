package controllers

import (
	"strconv"

	"gitee.com/masx200/to-do-list-go-sql-vue/backend/database"
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/models"
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
			var parameters models.QueryParameters
			err := c.ShouldBindQuery(&parameters)
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
			var limit = 100
			if parameters.Limit > 0 {
				limit = parameters.Limit
			}
			var page = 0
			if parameters.Page > 0 {
				page = parameters.Page
			}

			tdi, err := database.FindItems(db, limit, page, model, &query)
			if err != nil {
				c.String(500, err.Error())
			} else {
				c.JSON(200, tdi)
			}

		})
}
