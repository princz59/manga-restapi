package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"

	// "manga-restapi/configs"
	"manga-restapi/models"
)

func (ctrls *DBController) GetMangas(c *gin.Context) {
	limit_s := c.Query("limit")
	offset_s := c.Query("offset")

	limit, _ := strconv.Atoi(limit_s)
	offset, _ := strconv.Atoi(offset_s)

	sqlStatement := `
		SELECT * FROM directories 
		OFFSET $1
		LIMIT $2
	`

	rows, err := ctrls.Database.Query(sqlStatement, offset, limit)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	mangas := []models.MangaDirectories{}
	for rows.Next() {
		var manga models.MangaDirectories
		var genres_str string
		err = rows.Scan(&manga.ID, &manga.Name, &manga.Status, &genres_str, &manga.URL_PATH, &manga.COVER_URL_PATH)
		if err != nil {
			log.Fatalln(err)
		}
		manga.Genres = strings.Split(genres_str, ",")
		mangas = append(mangas, manga)
	}

	c.JSON(http.StatusOK, gin.H{"results": mangas})
}
