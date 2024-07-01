package app

import (
	"database/sql"
	"time"

	"github.com/raihanmd/dependency_injection/helper"

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
