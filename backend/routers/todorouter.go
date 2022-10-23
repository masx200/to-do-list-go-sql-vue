package routers

import (
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TodoRoute[T any](r *gin.Engine, db *gorm.DB, prefix string, model *T) {

	controllers.GETHEADItems(r, db, prefix, model)
	controllers.POSTItem(r, db, prefix, model)

	controllers.DELETEItem(r, db, prefix, model)
	controllers.PUTItem(r, db, prefix, model)
}
