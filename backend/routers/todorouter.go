package routers

import (
	"github.com/masx200/to-do-list-go-sql-vue/backend/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TodoRoute[T any](r *gin.Engine, GetDB func() (*gorm.DB, error), prefix string, model *T) {

	controllers.GETItems(r, GetDB, prefix, model)
	controllers.POSTItem(r, GetDB, prefix, model)

	controllers.DELETEItem(r, GetDB, prefix, model)
	controllers.PUTItem(r, GetDB, prefix, model)
	controllers.PATCHItem(r, GetDB, prefix, model)
}
