package main

import (
	"errors"
	"strconv"

	"gitee.com/masx200/to-do-list-go-sql-vue/backend/configs"
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/database"
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/models"
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/routers"
	"github.com/cuigh/auxo/util/lazy"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	var LoadConfig = configs.LoadConfig
	var model = &models.ToDoItem{}
	config := LoadConfig()

	var lazyDB = lazy.Value[*gorm.DB]{New: func() (db *gorm.DB, err error) {

		defer func() {

			if err := recover(); err != nil {

				e, o := err.(error)

				if o {
					err = e
				} else {
					err = errors.New("unknown error")
				}
			}
		}()
		db = database.ConnectDatabase(config.Dsn, &models.ToDoItem{}, "to_do_items", config.Debug)
		return
	}}

	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	var GetDB = func() (*gorm.DB, error) { return lazyDB.Get() }
	routers.TodoRoute(r, GetDB, "/todoitem", model)

	r.Run(":" + strconv.Itoa(config.Port))

}
