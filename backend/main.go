package main

import (
	"errors"
	"strconv"

	"github.com/cuigh/auxo/util/lazy"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/masx200/to-do-list-go-sql-vue/backend/configs"
	"github.com/masx200/to-do-list-go-sql-vue/backend/database"
	"github.com/masx200/to-do-list-go-sql-vue/backend/models"
	"github.com/masx200/to-do-list-go-sql-vue/backend/routers"
	"gorm.io/gorm"
)

func main() {
	var LoadConfig = configs.LoadConfig
	var model = &models.ToDoItem{}
	config := LoadConfig()

	var lazyDB = lazy.Value[*gorm.DB]{New: func() (db *gorm.DB, err error) {

		db, err = database.ConnectDatabase(config.Dsn, &models.ToDoItem{}, "to_do_items", config.Debug)
		return
	}}

	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	var GetDB = func() (*gorm.DB, error) {
		var db, err = lazyDB.Get()
		if err != nil {
			return nil, err
		}
		if db == nil {
			return nil, errors.New("db is  nil pointer")
		}
		return db, err
	}
	routers.TodoRoute(r, GetDB, "/todoitem", model)

	r.Run(":" + strconv.Itoa(config.Port))

}
