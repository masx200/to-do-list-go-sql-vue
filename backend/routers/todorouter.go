package routers

import (
	"strconv"

	"gitee.com/masx200/to-do-list-go-sql-vue/backend/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TodoRoute[T any](r *gin.Engine, db *gorm.DB, prefix string) {

	r.GET(prefix, func(c *gin.Context) {

		qsid := c.Query("id")

		if len(qsid) != 0 {
			c.Abort()
			id, err := strconv.Atoi(qsid)

			if err != nil {
				c.String(400, err.Error())
				return
			}
			var item = new(T)
			item, err = database.GetItem(db, item, uint(id))
			if err != nil {
				c.String(404, err.Error())
			} else {
				c.JSON(200, item)
			}
			return
		} else {
			c.Next()
		}

	}, func(c *gin.Context) {
		qsid := c.Query("id")

		if len(qsid) != 0 {
			return
		}

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
		tdi, err := database.FindItems(db, []T{}, limit, page)
		if err != nil {
			c.String(500, err.Error())
		} else {
			c.JSON(200, tdi)
		}
		// return
	})
	r.POST(prefix, func(c *gin.Context) {
		var item T
		err := c.ShouldBindJSON(&item)
		if err != nil {
			c.String(400, err.Error())
			return
		}
		err = database.CreateItem(db, &item)
		if err != nil {
			c.String(500, err.Error())
		} else {
			c.JSON(200, item)
		}
		// return
	})

	DELETEItem[T](r, db, prefix)
	PUTItem[T](r, db, prefix)
}

func DELETEItem[T any](r *gin.Engine, db *gorm.DB, prefix string) {
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
		item, err = database.GetItem(db, item, uint(id))
		/* 保持接口的幂等性 */
		if err != nil {
			c.JSON(200, gin.H{
				"id": id,
			})
			return
		}
		err = database.DeleteItem(db, new(T), uint(id))
		if err != nil {
			c.String(500, err.Error())
		} else {
			c.JSON(200, item)
		}
		// return
	})
}

func PUTItem[T any](r *gin.Engine, db *gorm.DB, prefix string) {
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
		item, err = database.GetItem(db, item, uint(id))
		if err != nil {
			c.String(404, err.Error())
		} else {
			c.JSON(200, item)
		}
		// return
	})
}
