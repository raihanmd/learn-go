package test

import (
	"context"
	"golang_database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	conn := golang_database.CreateConnection()
	defer conn.Close()

	conn.Exec("DELETE FROM users WHERE id = 'test'")

	m.Run()

	conn.Exec("DELETE FROM users WHERE id = 'test'")
}

func TestInsert(t *testing.T) {
	conn := golang_database.CreateConnection()
	defer conn.Close()

	_, err := conn.ExecContext(context.Background(), "INSERT INTO users(id, name) VALUES('test', 'test')")
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
