package controllers

import (
	"context"
	"errors"

	"gitee.com/masx200/to-do-list-go-sql-vue/backend/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PUTItem[T any](r *gin.Engine, db *gorm.DB, prefix string, model *T) {
	r.PUT(prefix, func(c *gin.Context) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		var inputs []map[string]any
		err := c.ShouldBindJSON(&inputs)
		if err != nil {
			c.String(400, err.Error())
			return
		}
		var ch = make(chan TWO[map[string]any, error])
		var output = func(res map[string]any, err error) {
			select {
			case <-ctx.Done():
				{
					return
				}
			case ch <- TWO[map[string]any, error]{res, err}:
				{
					return
				}
			}
		}
		// 开始事务
		tx := CloneDB(db).Begin()
		for _, input := range inputs {
			qsid, o := input["id"]
			if !o {
				// 遇到错误时回滚事务
				tx.Rollback()
				c.String(400, "expect id but not found")
				return
			}
			id, ok := qsid.(float64)

			if !ok {
				// 遇到错误时回滚事务
				tx.Rollback()
				c.String(400, err.Error())
				return
			}
			var item = input
			go func(id float64, item map[string]any) {
				defer func() {

					if err := recover(); err != nil {

						e, o := err.(error)

						if o {
							output(nil, e)
						} else {
							output(nil, errors.New("unknown error"))
						}
					}
				}()
				/* put 先删除再创建修改 */

				err = database.UpsertItem(tx, model, item, uint(id))
				/* 保持接口的幂等性 */
				if err != nil {

					output(nil, err)

					return
				}
				res, err := database.GetItem(tx, model, uint(id))
				if err != nil {
					output(nil, err)

					return

				} else {
					output(res, nil)

					return

				}
			}(id, item)
		}

		var results = []map[string]interface{}{}
		for range inputs {
			two := <-ch
			res := two.First
			err := two.Second
			if err != nil {
				// 遇到错误时回滚事务
				tx.Rollback()
				c.String(500, err.Error())
				return
			}
			results = append(results, res)
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
