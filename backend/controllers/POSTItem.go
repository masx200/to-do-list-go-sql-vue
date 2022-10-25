package controllers

import (
	"context"
	"errors"

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
		for _, input := range inputs {
			delete(input, "id")
			go func(input map[string]any) {
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

				var item = database.MapToStruct[T](input)
				id, err := database.CreateItem(createDB, model, item)
				if err != nil {
					output(nil, err)
					return
				} else {

					res, err := database.GetItem(createDB, model, uint(id))
					if err != nil {
						output(nil, err)
						return
					} else {
						output(res, nil)

						return
					}

				}
			}(input)
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
