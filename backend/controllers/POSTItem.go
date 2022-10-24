package controllers

import (
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TWO[T any, Y any] struct {
	First  T
	Second Y
}

func POSTItem[T any](r *gin.Engine, createDB func() *gorm.DB, prefix string, model *T) {
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
		var ch = make(chan TWO[map[string]any, error])

		for _, input := range inputs {

			go func(input map[string]any, ch chan TWO[map[string]any, error]) {

				var item = database.MapToStruct[T](input)
				id, err := database.CreateItem(createDB, model, item)
				if err != nil {
					ch <- TWO[map[string]any, error]{nil, err}
					return
				} else {

					res, err := database.GetItem(createDB, model, uint(id))
					if err != nil {
						ch <- TWO[map[string]any, error]{nil, err}
						return
					} else {
						ch <- TWO[map[string]any, error]{res, nil}
						return
					}

				}
			}(input, ch)
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
