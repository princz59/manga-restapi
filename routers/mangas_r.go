package routers

import (
	"manga-restapi/controllers"
	"github.com/gin-gonic/gin"
)

func SetMangasRoutes(router *gin.RouterGroup, ctrls *controllers.DBController) {
	router.GET("mangas", ctrls.GetMangas)
}