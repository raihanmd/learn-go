package app

import (
	"database/sql"
	"restful_api/helper"
	"time"

	_ "modernc.org/sqlite"
)

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite", "database.db")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
