package test

import (
	"context"
	"database/sql"
	"golang_database"
	"golang_database/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	conn := golang_database.CreateConnection()
	defer conn.Close()

	conn.Exec("DELETE FROM users WHERE id IN ('test', 'new')")
	conn.Exec("DELETE FROM fruits")

	m.Run()

	conn.Exec("DELETE FROM users WHERE id IN ('test', 'new')")
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

	conn := golang_database.CreateConnection()
	defer conn.Close()

	rows, err := conn.QueryContext(context.Background(), "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM users WHERE name != 'test' ORDER BY id ASC")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	expect := []model.User{
		{
			Id:        "admin",
			Name:      "admin",
			Email:     sql.NullString{Valid: false},
			Balance:   0,
			Rating:    0,
			BirthDate: sql.NullTime{Valid: false},
			Married:   false,
			Created:   time.Date(2024, 6, 14, 19, 27, 26, 0, time.UTC)},
		{
			Id:        "budi",
			Name:      "Budi",
			Email:     sql.NullString{Valid: false},
			Balance:   50000,
			Rating:    2.5,
			BirthDate: sql.NullTime{Time: time.Date(2024, 6, 14, 0, 0, 0, 0, time.UTC), Valid: true},
			Married:   false,
			Created:   time.Date(2024, 6, 14, 17, 7, 1, 0, time.UTC)},
		{
			Id:        "eko",
			Name:      "Eko",
			Email:     sql.NullString{String: "eko@emial.co.id", Valid: true},
			Balance:   100000,
			Rating:    3.4,
			BirthDate: sql.NullTime{Time: time.Date(2024, 6, 14, 0, 0, 0, 0, time.UTC), Valid: true},
			Married:   false,
			Created:   time.Date(2024, 6, 14, 17, 6, 35, 0, time.UTC)},
	}

	var actual = []model.User{}

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

		actual = append(actual, model.User{
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
