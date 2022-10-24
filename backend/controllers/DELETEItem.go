package controllers

import (
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DELETEItem[T any](r *gin.Engine, createDB func() *gorm.DB, prefix string, model *T) {
	r.DELETE(prefix, func(c *gin.Context) {

		var inputs []map[string]any
		err := c.ShouldBindJSON(&inputs)
		if err != nil {
			c.String(400, err.Error())
			return
		}
		var ch = make(chan TWO[map[string]any, error])
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

			go func(id float64, ch chan TWO[map[string]any, error]) {

				var item = new(T)
				res, err := database.GetItem(createDB, item, uint(id))
				/* 保持接口的幂等性 */
				if err != nil {

					ch <- TWO[map[string]any, error]{gin.H{
						"id": id,
					}, nil}

					return
				}
				err = database.DeleteItem(createDB, new(T), uint(id))
				if err != nil {
					ch <- TWO[map[string]any, error]{nil, err}
					return

				} else {
					ch <- TWO[map[string]any, error]{res, nil}
					return

				}
			}(id, ch)
		}

		var results = []map[string]interface{}{}
		for range inputs {
			two := <-ch
			res := two.First
			err := two.Second
			if err != nil {
				c.String(500, err.Error())
				return
			}
			results = append(results, res)
		}
		c.JSON(200, results)

	})
}
