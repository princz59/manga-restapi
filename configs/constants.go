package configs

import (
	"os"
	"log"
	"strconv"
	"github.com/joho/godotenv"
)

var (
	BASE_URL 		   string = ""
	PORT   			   string = ""
	DB_NAME            string = ""
	DB_HOST            string = ""
	DB_USER            string = ""
	DB_PASSWORD		   string = ""
	DB_PORT_STR        string = ""
	DB_PORT 		   int    = 0
)


func InitConstantVariable() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	DB_HOST = os.Getenv("DB_HOST")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_PORT, _ = strconv.Atoi((os.Getenv("DB_PORT")))
	PORT = os.Getenv("PORT")
	BASE_URL = os.Getenv("BASE_URL")
}