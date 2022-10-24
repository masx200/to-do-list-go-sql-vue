package routers

import (
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TodoRoute[T any](r *gin.Engine, createDB func() *gorm.DB, prefix string, model *T) {

	controllers.GETItems(r, createDB, prefix, model)
	controllers.POSTItem(r, createDB, prefix, model)

	controllers.DELETEItem(r, createDB, prefix, model)
	controllers.PUTItem(r, createDB, prefix, model)
}
