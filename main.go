package main

import (
	"log"
	"net/http"

	"manga-restapi/configs"
	"manga-restapi/routers"
	database "manga-restapi/configs/db"
	
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	configs.InitConstantVariable()
	db, _ := database.InitDatabase()
	db.Ping()
	router = gin.Default()
	routers.InitRouters(router, db)
}


func main() {
	log.Println("Server Running on Port: ", configs.PORT)
	http.ListenAndServe(":" + configs.PORT, router)
}