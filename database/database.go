package golang_database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection() *sql.DB {
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang")
	if err != nil {
		panic(err)
	}
	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(100)
	conn.SetConnMaxIdleTime(5 * time.Minute)
	conn.SetConnMaxLifetime(60 * time.Second)
	return conn
}
