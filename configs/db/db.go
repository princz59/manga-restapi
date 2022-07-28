package database

import (
	"manga-restapi/configs"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDatabase() (db *sql.DB, err error){

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		configs.DB_HOST, configs.DB_PORT, configs.DB_USER, configs.DB_PASSWORD, configs.DB_NAME)
	
		db, err = sql.Open("postgres", psqlInfo)

	return
}	