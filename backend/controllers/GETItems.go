package controllers

import (
	"strconv"

	"gitee.com/masx200/to-do-list-go-sql-vue/backend/database"
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GETItems[T any](r *gin.Engine, createDB func() *gorm.DB, prefix string, model *T) {
	r.GET(prefix, func(c *gin.Context) {

		qsid := c.Query("id")

		if len(qsid) != 0 {
			c.Abort()
			id, err := strconv.Atoi(qsid)

			if err != nil {
				c.String(400, err.Error())
				return
			}

			res, err := database.GetItem(createDB, model, uint(id))
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

			var item T

			err = c.ShouldBindQuery(&item)
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
			var order = "id"
			if len(parameters.Order) > 0 {
				order = parameters.Order
			}
			var direction = "asc"
			if len(parameters.Direction) > 0 {
				direction = parameters.Direction
			}
			var qm = database.StructToMap(&item)

			values := c.Request.URL.Query()
			var query = map[string]any{}
			for k := range values {
				if v, o := qm[k]; o {
					query[k] = v
				}

			}
			delete(query, "id")
			tdi, err := database.FindItems(createDB, limit, page, model, query, order, direction)
			if err != nil {
				c.String(500, err.Error())
			} else {
				c.JSON(200, tdi)
			}

		})
}
