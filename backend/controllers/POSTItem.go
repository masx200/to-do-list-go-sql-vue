package controllers

import (
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/database"
	"github.com/akrennmair/slice"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TWO[T any, Y any] struct {
	First  T
	Second Y
}

func POSTItem[T any](r *gin.Engine, db *gorm.DB, prefix string, model *T) {
	r.POST(prefix, func(c *gin.Context) {

		var inputs []map[string]any
		err := c.ShouldBindJSON(&inputs)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		for _, input := range inputs {
			delete(input, "id")

		}

		var items = slice.Map(inputs, func(input map[string]any) *T {

			return database.MapToStruct[T](input)
		})

		// 开始事务
		tx := CloneDB(db).Begin()
		ids, err := database.CreateItems(tx, model, items)
		if err != nil {
			// 遇到错误时回滚事务
			tx.Rollback()
			c.String(500, err.Error())
			return
		}
		var results []map[string]interface{}
		results, err = database.FindByIDs(tx, model, ids)

		if err != nil {

			// 遇到错误时回滚事务
			tx.Rollback()
			c.String(500, err.Error())
			return
		}
		// 否则，提交事务
		err = tx.Commit().Error
		if err == nil {
			c.JSON(200, results)
		} else {
			c.String(500, err.Error())
			return
		}
	})
}
