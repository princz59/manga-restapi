package controllers

import (
	"database/sql"
)

type DBController struct {
	Database *sql.DB
}