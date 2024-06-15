package model

import (
	"database/sql"
	"time"
)

// ? Use sql.Null... for NULLABLE Data Type DB
type User struct {
	Id        string
	Name      string
	Password  string
	Email     sql.NullString
	Balance   int32
	Rating    float64
	BirthDate sql.NullTime
	Married   bool
	Created   time.Time
}
