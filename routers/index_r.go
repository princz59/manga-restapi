package routers

import (
	"database/sql"
	"manga-restapi/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine, db *sql.DB) {
	
	ctrls := &controllers.DBController{Database: db}
	api := r.Group("/api")

	SetMangasRoutes(api, ctrls)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Manga Reader API!",
		})
	})
}