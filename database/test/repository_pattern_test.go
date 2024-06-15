package test

import (
	"context"
	"database/sql"
	"golang_database"
	"golang_database/model"
	"golang_database/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertRepo(t *testing.T) {
	userRepository := repository.NewUserRepository(golang_database.CreateConnection())

	newUser := model.User{
		Id:        "new",
		Name:      "new",
		Password:  "new",
		Email:     sql.NullString{String: "new", Valid: true},
		Balance:   10000,
		Rating:    10,
		BirthDate: sql.NullTime{Valid: false},
		Married:   true,
	}

	user, err := userRepository.Insert(context.Background(), newUser)
	if err != nil {
		panic(err)
	}
	t.Log(user)
	assert.Equal(t, newUser, user)
}

func TestFindByIdRepo(t *testing.T) {
	TestInsertRepo(&testing.T{})

	userRepository := repository.NewUserRepository(golang_database.CreateConnection())

	user, err := userRepository.FindById(context.Background(), "new")
	if err != nil {
		panic(err)
	}

	t.Log(user)
	assert.Equal(t, "new", user.Id)
}

func TestFindAllRepo(t *testing.T) {
	userRepository := repository.NewUserRepository(golang_database.CreateConnection())

	users, err := userRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	t.Log(users)
	assert.NotNil(t, users)
}
