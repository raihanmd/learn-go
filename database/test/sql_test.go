package test

import (
	"context"
	"database/sql"
	"golang_database"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	conn := golang_database.CreateConnection()
	defer conn.Close()

	conn.Exec("DELETE FROM users WHERE id = 'test'")
	conn.Exec("DELETE FROM fruits")

	m.Run()

	conn.Exec("DELETE FROM users WHERE id = 'test'")
}

func TestInsert(t *testing.T) {
	conn := golang_database.CreateConnection()
	defer conn.Close()

	res, err := conn.ExecContext(context.Background(), "INSERT INTO users(id, name, password) VALUES('test', 'test', 'test')")
	if err != nil {
		panic(err)
	}

	// ? Get Last Inserted ID For AutoIncrement Id
	t.Log(res.LastInsertId())

	row, err := conn.QueryContext(context.Background(), "SELECT id, name FROM users WHERE id = 'test'")
	if err != nil {
		panic(err)
	}

	for row.Next() {
		var id, name string

		err := row.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, "test", id)
		assert.Equal(t, "test", name)
	}

	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	conn := golang_database.CreateConnection()
	defer conn.Close()

	_, err := conn.ExecContext(context.Background(), "UPDATE users SET name = 'test_update' WHERE id = 'test'")
	if err != nil {
		panic(err)
	}

	row, err := conn.QueryContext(context.Background(), "SELECT id, name FROM users WHERE id = 'test'")
	if err != nil {
		panic(err)
	}

	for row.Next() {
		var id, name string

		err := row.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, "test", id)
		assert.Equal(t, "test_update", name)
	}

	assert.Nil(t, err)
}

func TestMultiDataType(t *testing.T) {
	// ? Use sql.Null... for NULLABLE Data Type DB
	type User struct {
		Id        string
		Name      string
		Email     sql.NullString
		Balance   int32
		Rating    float64
		BirthDate sql.NullTime
		Married   bool
		Created   time.Time
	}

	conn := golang_database.CreateConnection()
	defer conn.Close()

	rows, err := conn.QueryContext(context.Background(), "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM users WHERE name != 'test' ORDER BY id ASC")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	expect := []User{
		{"admin", "admin", sql.NullString{Valid: false}, 0, 0, sql.NullTime{Valid: false}, false, time.Date(2024, 6, 14, 19, 27, 26, 0, time.UTC)},
		{"budi", "Budi", sql.NullString{Valid: false}, 50000, 2.5, sql.NullTime{Time: time.Date(2024, 6, 14, 0, 0, 0, 0, time.UTC), Valid: true}, false, time.Date(2024, 6, 14, 17, 7, 1, 0, time.UTC)},
		{"eko", "Eko", sql.NullString{String: "eko@emial.co.id", Valid: true}, 100000, 3.4, sql.NullTime{Time: time.Date(2024, 6, 14, 0, 0, 0, 0, time.UTC), Valid: true}, false, time.Date(2024, 6, 14, 17, 6, 35, 0, time.UTC)},
	}

	var actual = []User{}

	for rows.Next() {
		var (
			id, name  string
			email     sql.NullString
			balance   int32
			rating    float64
			birthDate sql.NullTime
			createdAt time.Time
			married   bool
		)

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		actual = append(actual, User{
			Id:        id,
			Name:      name,
			Email:     email,
			Balance:   balance,
			Rating:    rating,
			BirthDate: birthDate,
			Married:   married,
			Created:   createdAt,
		})

	}

	for _, user := range actual {
		t.Log("\nId:", user.Id, "\nName:", user.Name, "\nEmail:", user.Email, "\nBalance:", user.Balance, "\nRating:", user.Rating, "\nBirthDate:", user.BirthDate, "\nMarried:", user.Married, "\nCreated:", user.Created)
	}

	assert.Equal(t, expect, actual)
}

func TestSqlInjection(t *testing.T) {
	conn := golang_database.CreateConnection()
	defer conn.Close()

	username := "admin"
	password := "admin"

	row, err := conn.QueryContext(context.Background(), "SELECT name from users WHERE name=? AND password=? LIMIT 1", username, password)
	if err != nil {
		panic(err)
	}
	if row.Next() {
		t.Log("Login Success")
	} else {
		t.Log("Login Failed")
	}

	assert.True(t, row.Next())
}

func TestPrepareStatement(t *testing.T) {
	conn := golang_database.CreateConnection()
	defer conn.Close()

	stmt, err := conn.PrepareContext(context.Background(), "INSERT INTO fruits (name) VALUES (?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	fruits := []string{"apple", "banana", "melon", "orange", "grape"}

	for _, fruit := range fruits {
		res, err := stmt.ExecContext(context.Background(), fruit)
		if err != nil {
			panic(err)
		}
		lastId, _ := res.LastInsertId()
		t.Log("Fruit Id:", lastId)
		assert.Nil(t, err)
	}
}

func TestTransaction(t *testing.T) {
	conn := golang_database.CreateConnection()
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		panic(err)
	}

	_, err = tx.ExecContext(context.Background(), "INSERT INTO fruits (name) VALUES (?)", "avocado")
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
