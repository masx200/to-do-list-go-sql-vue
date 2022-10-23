package routers

import (
	"gitee.com/masx200/to-do-list-go-sql-vue/backend/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TodoRoute[T any](r *gin.Engine, db *gorm.DB, prefix string) {

	controllers.GETItems[T](r, db, prefix)
	controllers.POSTItem[T](r, db, prefix)

	controllers.DELETEItem[T](r, db, prefix)
	controllers.PUTItem[T](r, db, prefix)
}
